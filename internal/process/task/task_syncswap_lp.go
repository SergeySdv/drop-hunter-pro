package task

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/pkg/errors"
)

type SyncSwapLPTask struct {
	cancel func()
}

func (t *SyncSwapLPTask) Stop() error {
	t.cancel()
	return nil
}

func (t *SyncSwapLPTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *SyncSwapLPTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_SyncSwapLPTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_SyncSwapLPTask) call an ambulance!")
	}

	p := l.SyncSwapLPTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning:

		task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	client, _, err := NewZkSyncClient(profile, p.Network)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		estimation, err := EstimateSyncSwapLPCost(ctx, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateSyncSwapLPCost")
		}
		res, gas, err := SyncSwapLP(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "SyncSwapLP")
		}

		p.Tx = NewTx(res.SwapTx, gas)
		if err := a.AddTx(ctx, res.ApproveATxHash); err != nil {
			return nil, err
		}
		if err := a.AddTx(ctx, res.ApproveBTxHash); err != nil {
			return nil, err
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, client, a); err != nil {
		return nil, err
	}
	if err := a.AddTx2(ctx, p.Tx); err != nil {
		return nil, err
	}

	if p.GetTx().GetTxCompleted() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func SyncSwapLP(ctx context.Context, profile *halp.Profile, p *v1.SyncSwapLPTask, client zksyncera.LP, estimation *v1.EstimationTx) (*zksyncera.SyncSwapLiquidityRes, *defi.Gas, error) {

	var err error
	if client == nil {
		client, _, err = NewZkSyncClient(profile, p.Network)
		if err != nil {
			return nil, nil, err
		}
	}

	wallet, err := zksyncera.NewWalletTransactor(profile.WalletPK, client.GetNetworkId())
	if err != nil {
		return nil, nil, err
	}

	var am, b *big.Int

	if p.Add {
		balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: wallet.WalletAddr,
			Token:         p.A,
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetBalance")
		}
		b = balance.WEI
		am, err = defi.ResolveAmount(p.Amount, balance.WEI)
		if err != nil {
			return nil, nil, err
		}
	} else {
		poolAddr, err := client.GetSyncSwapPool(ctx, &zksyncera.GetSyncSwapPoolReq{A: p.A, B: p.B})
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.GetSyncSwapPool")
		}
		b, err = client.SyncSwapLPBalance(ctx, *poolAddr, wallet.WalletAddr)
		if err != nil {
			return nil, nil, errors.Wrap(err, "client.SyncSwapLPBalance")
		}
		am, err = defi.ResolveAmount(p.Amount, b)
		if err != nil {
			return nil, nil, err
		}
	}

	if am == nil || am.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("not enough balance of " + p.A.String())
	}

	estimateOnly := estimation == nil
	var Gas *defi.Gas
	if estimateOnly {
		am = defi.Percent(am, 90)
		Gas = nil
	} else {
		gas, err := GasManager(estimation, profile.Settings, p.Network)
		if err != nil {
			return nil, nil, err
		}
		Gas = gas
	}

	res, err := client.SyncSwapLiquidity(ctx, &zksyncera.SyncSwapLiquidityReq{
		Amount:       am,
		A:            p.A,
		B:            p.B,
		WalletPK:     wallet.PK,
		Add:          p.Add,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func EstimateSyncSwapLPCost(ctx context.Context, profile *halp.Profile, p *v1.SyncSwapLPTask, client zksyncera.LP) (*v1.EstimationTx, error) {
	res, _, err := SyncSwapLP(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.EstimatedGasCost, p.Network), nil
}
