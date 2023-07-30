package task

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WithdrawExchange struct {
	cancel func()
}

func (t *WithdrawExchange) SetTimeout(td time.Duration) {

}

func (t *WithdrawExchange) Stop() error {
	t.cancel()
	return nil
}

func (t *WithdrawExchange) Reset(ctx context.Context, a *Input) error {
	task := a.Task
	l, ok := task.Task.Task.(*v1.Task_WithdrawExchangeTask)
	if !ok {
		return errors.New("a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask) call an ambulance!")
	}
	p := l.WithdrawExchangeTask

	p.TxId = nil
	p.WithdrawOrderId = nil

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *WithdrawExchange) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_WithdrawExchangeTask) call an ambulance!")
	}
	p := l.WithdrawExchangeTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRunning:
		if p.WithdrawOrderId == nil {
			task.Status = v1.ProcessStatus_StatusReady
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, errors.Wrap(err, "a.UpdateTask")
			}
		}

		withdrawer, err := a.WithdrawerRepository.GetWithdrawer(ctx, p.WithdrawerId, a.UserId)
		if err != nil {
			return nil, errors.Wrap(err, "a.WithdrawerRepository.GetWithdrawer")
		}

		profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
		if err != nil {
			return nil, errors.Wrap(err, "a.ProfileRepository.GetProfile")
		}

		exchangeWithdrawer, err := uniclient.NewExchangeWithdrawer(withdrawer, profile)
		if err != nil {
			return nil, errors.Wrap(err, "uniclient.NewExchangeWithdrawer")
		}

		txId, err := exchangeWithdrawer.WaitConfirm(taskContext, *p.WithdrawOrderId)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return a.Task, nil
			}
			return nil, errors.Wrap(err, "exchangeWithdrawer.WaitConfirm")
		}

		p.TxId = txId
		after := v1.ProcessStatus_StatusDone
		task.Status = after
		task.Task.Description = p.String()
		task.FinishedAt = timestamppb.Now()
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
		return task, nil

	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:

		if p.WithdrawOrderId != nil {
			after := v1.ProcessStatus_StatusRunning
			a.Task.Status = after
			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, errors.Wrap(err, "a.UpdateTask")
			}
			return task, nil
		}

		after := v1.ProcessStatus_StatusRunning
		a.Task.Status = after
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
	}

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, errors.Wrap(err, "ProfileRepository.GetProfile")
	}

	withdrawer, err := a.WithdrawerRepository.GetWithdrawer(ctx, p.WithdrawerId, a.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "WithdrawerRepository.GetWithdrawer")
	}

	exchangeWithdrawer, err := uniclient.NewExchangeWithdrawer(withdrawer, profile)
	if err != nil {
		return nil, errors.Wrap(err, "uniclient.NewExchangeWithdrawer")
	}

	transactor, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		return nil, errors.Wrap(err, "defi.NewWalletTransactor")
	}

	if p.UseExternalAddr != nil && *p.UseExternalAddr {

	} else {
		p.WithdrawAddr = &transactor.WalletAddrHR
	}

	if p.SendAllCoins != nil && *p.SendAllCoins {
		b, err := exchangeWithdrawer.GetBalance(ctx, p.Token)
		if err != nil {
			return nil, errors.Wrap(err, "exchangeWithdrawer.GetBalance")
		}
		amount := lib.FloatToString(b)
		p.Amount = &amount
	} else {
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
	}

	res, err := exchangeWithdrawer.Withdraw(taskContext, &exchange.WithdrawRequest{
		ToAddress: *p.WithdrawAddr,
		Amount:    *p.Amount,
		Network:   p.Network,
		Token:     p.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "exchangeWithdrawer.Withdraw")
	}

	p.WithdrawOrderId = &res.WithdrawId

	if err := a.UpdateTask(ctx, task); err != nil {
		return nil, errors.Wrap(err, "UpdateTask")
	}
	return task, nil
}
