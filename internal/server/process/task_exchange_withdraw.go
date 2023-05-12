package process

import (
	"context"
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/exchange"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/uniclient"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WithdrawExchange struct {
}

func (t *WithdrawExchange) Run(ctx context.Context, a *TaskArg) (*TaskRes, error) {

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask) call an ambulance!")
	}
	p := l.WithdrawExchangeTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return &TaskRes{Task: a.Task}, nil
	case v1.ProcessStatus_StatusRunning:
		if p.WithdrawOrderId == nil {
			before := a.Task.Status
			after := v1.ProcessStatus_StatusReady
			a.Task.Status = after
			if err := a.Rep.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
				return nil, err
			}
			_ = a.Rep.RecordStatusChanged(ctx, task.Id, before, after)
		}

		withdrawer, err := a.Rep.GetWithdrawer(ctx, p.WithdrawerId)
		if err != nil {
			return nil, err
		}

		exchangeWithdrawer, err := uniclient.NewExchangeWithdrawer(withdrawer)
		if err != nil {
			return nil, err
		}

		txId, err := exchangeWithdrawer.WaitConfirm(ctx, *p.WithdrawOrderId)
		if err != nil {
			return nil, err
		}

		p.TxId = txId
		before := a.Task.Status
		after := v1.ProcessStatus_StatusDone
		task.Status = after
		task.Task.Description = p.String()
		task.FinishedAt = timestamppb.Now()
		if err := a.Rep.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, err
		}
		_ = a.Rep.RecordStatusChanged(ctx, task.Id, before, after)
		return &TaskRes{
			Task: task,
		}, nil

	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:

		if p.WithdrawOrderId != nil {
			before := a.Task.Status
			after := v1.ProcessStatus_StatusRunning
			a.Task.Status = after
			if err := a.Rep.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
				return nil, err
			}
			_ = a.Rep.RecordStatusChanged(ctx, task.Id, before, after)
			return &TaskRes{
				Task: task,
			}, nil
		}

		before := a.Task.Status
		after := v1.ProcessStatus_StatusRunning
		a.Task.Status = after
		if err := a.Rep.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, err
		}
		_ = a.Rep.RecordStatusChanged(ctx, task.Id, before, after)
	}

	profile, err := a.Rep.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	withdrawer, err := a.Rep.GetWithdrawer(ctx, p.WithdrawerId)
	if err != nil {
		return nil, err
	}

	exchangeWithdrawer, err := uniclient.NewExchangeWithdrawer(withdrawer)
	if err != nil {
		return nil, err
	}

	transactor, err := defi.NewWalletTransactor(profile.MmskPk)
	if err != nil {
		return nil, err
	}

	min, err := lib.StringToFloat(p.AmountMin)
	if err != nil {
		return nil, err
	}
	max, err := lib.StringToFloat(p.AmountMax)
	if err != nil {
		return nil, err
	}

	am := lib.FloatToString(lib.RandFloatRange(min, max))
	p.Amount = &am

	res, err := exchangeWithdrawer.Withdraw(ctx, &exchange.WithdrawRequest{
		ToAddress: transactor.WalletAddrHR,
		Amount:    *p.Amount,
		Network:   p.Network,
		Token:     p.Token,
	})
	if err != nil {
		return nil, err
	}

	p.WithdrawOrderId = &res.WithdrawId
	task.Task.Description = p.String()

	if err := a.Rep.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
		return nil, err
	}

	_ = a.Rep.RecordStatusChanged(ctx, task.Id, v1.ProcessStatus_StatusRunning, v1.ProcessStatus_StatusDone)

	return &TaskRes{
		Task: task,
	}, nil
}
