package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type WethTask struct {
	cancel func()
}

func (t *WethTask) Stop() error {
	t.cancel()
	return nil
}

func (t *WethTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

var ErrWethTaskError = errors.New("weith task error")

func (t *WethTask) Run(ctx context.Context, a *Input) (_ *v1.ProcessTask, err error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_WETHTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}

	p := l.WETHTask

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
		return nil, errors.Wrap(err, "a.Halper.Profile")
	}

	token := v1.Token_WETH
	if p.Wrap {
		token = v1.Token_ETH
	}

	client, err := uniclient.NewWETH(p.Network, profile.BaseConfig(p.Network))
	if err != nil {
		return nil, errors.Wrap(err, "uniclient.NewWETH")
	}

	if p.GetTx().GetTxId() == "" {
		b, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.WalletAddr,
			Token:         token,
		})
		if err != nil {
			return nil, errors.Wrap(err, "client.GetBalance")
		}

		am, err := defi.ResolveAmount(p.Amount, b.WEI)
		if err != nil {
			return nil, errors.Wrap(err, "defi.ResolveAmount")
		}

		estimation, err := EstimateWethTaskCost(ctx, p, profile)
		if err != nil {
			return nil, errors.Wrap(err, "EstimateWethTaskCost")
		}
		gas, err := GasManager(estimation, profile.Settings, p.Network)
		if err != nil {
			return nil, err
		}

		if token == client.GetNetworkToken() {
			am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
		}

		res, err := client.SwapWETH(ctx, &defi.WETHReq{
			Amount:       am,
			Wrap:         p.Wrap,
			WalletPK:     profile.WalletPK,
			EstimateOnly: false,
			Gas:          gas,
		})
		if err != nil {
			return nil, errors.Wrap(err, "client.SwapWETH")
		}

		p.Tx = NewTx(res.Tx, gas)
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, client, a); err != nil {
		return nil, errors.Wrap(err, "WaitTxComplete")
	}

	if err := a.AddTx2(ctx, p.Tx); err != nil {
		return nil, errors.Wrap(err, "a.AddTx2")
	}

	if p.GetTx().GetTxCompleted() {
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, errors.Wrap(err, "a.UpdateTask")
		}
	}

	return task, nil
}

func EstimateWethTaskCost(ctx context.Context, p *v1.WETHTask, profile *halp.Profile) (*v1.EstimationTx, error) {

	swapper, err := uniclient.NewWETH(p.Network, profile.BaseConfig(p.Network))
	if err != nil {
		return nil, err
	}

	token := v1.Token_WETH
	if p.Wrap {
		token = v1.Token_ETH
	}

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.WalletAddr,
		Token:         token,
	})
	if err != nil {
		return nil, err
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	swap, err := swapper.SwapWETH(ctx, &defi.WETHReq{
		Amount:       defi.Percent(am, 50),
		Wrap:         p.Wrap,
		WalletPK:     profile.WalletPK,
		EstimateOnly: true,
		Gas:          nil,
	})
	if err != nil {
		return nil, err
	}

	if token == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.ECost.TotalGasWei, am)
	}

	return GasStation(swap.ECost, p.Network), nil
}
