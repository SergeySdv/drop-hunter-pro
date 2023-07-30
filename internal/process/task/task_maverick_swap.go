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

type MaverickSwapTask struct {
	cancel func()
}

func (t *MaverickSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *MaverickSwapTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *MaverickSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_MaverickSwapTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_MaverickSwapTask) call an ambulance!")
	}

	p := l.MaverickSwapTask

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

		estimation, err := EstimateMaverickSwapCost(ctx, profile, p, client)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateSyncSwapLPCost")
		}
		res, gas, err := MaverickSwap(taskContext, profile, p, client, estimation)
		if err != nil {
			return nil, errors.Wrap(err, "SyncSwapLP")
		}

		p.Tx = NewTx(res.SwapTx, gas)
		if err := a.AddTx(ctx, res.ApproveTxHash); err != nil {
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

func MaverickSwap(ctx context.Context, profile *halp.Profile, p *v1.MaverickSwapTask, client zksyncera.MaverickSwapper, estimation *v1.EstimationTx) (*defi.DefaultSwapRes, *defi.Gas, error) {

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

	balance, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: wallet.WalletAddr,
		Token:         p.FromToken,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "client.GetBalance")
	}

	am, err := defi.ResolveAmount(p.Amount, balance.WEI)
	if err != nil {
		return nil, nil, err
	}

	if am == nil || am.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("not enough balance of " + p.FromToken.String())
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

	res, err := client.MaverickSwap(ctx, &defi.DefaultSwapReq{
		Network:      v1.Network_ZKSYNCERA,
		Amount:       am,
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     wallet.PK,
		EstimateOnly: estimateOnly,
		Gas:          Gas,
		Debug:        false,
	})
	if err != nil {
		return nil, nil, err
	}
	return res, Gas, nil
}

func EstimateMaverickSwapCost(ctx context.Context, profile *halp.Profile, p *v1.MaverickSwapTask, client zksyncera.MaverickSwapper) (*v1.EstimationTx, error) {
	res, _, err := MaverickSwap(ctx, profile, p, client, nil)
	if err != nil {
		return nil, err
	}

	return GasStation(res.EstimatedGasCost, p.Network), nil
}
