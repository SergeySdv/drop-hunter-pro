package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

type MuteioSwapTask struct {
	cancel func()
}

func (t *MuteioSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *MuteioSwapTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *MuteioSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_MuteioSwapTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_MuteioSwapTask) call an ambulance!")
	}

	p := l.MuteioSwapTask

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

	bc := profile.BaseConfig(p.Network)
	proxy, err := socks5.NewSock5ProxyString(bc.ProxyString, bc.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	client, err := zksyncera.NewMainNetClient(&zksyncera.ClientConfig{
		HttpCli:     proxy.Cli,
		RPCEndpoint: bc.RPCEndpoint,
	})
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		b, err := client.GetBalance(ctx, &defi.GetBalanceReq{
			WalletAddress: profile.WalletAddr,
			Token:         p.FromToken,
		})
		if err != nil {
			return nil, err
		}

		if b.Float == 0 {
			return nil, errors.New("not enough balance in " + p.FromToken.String())
		}

		am, err := defi.ResolveAmount(p.Amount, b.WEI)
		if err != nil {
			return nil, err
		}

		estimate, err := EstimateMuteioSwapCost(taskContext, profile, p)
		if err != nil {
			return nil, err
		}

		gas, err := GasManager(estimate, profile.Settings, p.Network)
		if err != nil {
			return nil, err
		}

		if p.FromToken == client.GetNetworkToken() {
			am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
		}

		res, err := client.MuteIOSwap(taskContext, &zksyncera.MuteIOSwapReq{
			Gas:       gas,
			Amount:    am,
			FromToken: p.FromToken,
			ToToken:   p.ToToken,
			WalletPK:  profile.WalletPK,
		})
		if err != nil {
			return nil, err
		}

		if err := a.AddTx(ctx, res.Allowance); err != nil {
			return nil, err
		}

		p.Tx = NewTx(res.Tx, gas)
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

func EstimateMuteioSwapCost(ctx context.Context, profile *halp.Profile, p *v1.MuteioSwapTask) (*v1.EstimationTx, error) {

	bc := profile.BaseConfig(p.Network)
	proxy, err := socks5.NewSock5ProxyString(bc.ProxyString, bc.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	swapper, err := zksyncera.NewMainNetClient(&zksyncera.ClientConfig{
		HttpCli:     proxy.Cli,
		RPCEndpoint: bc.RPCEndpoint,
	})
	if err != nil {
		return nil, err
	}
	wallet, err := zksyncera.NewWalletTransactor(profile.WalletPK, swapper.GetNetworkId())
	if err != nil {
		return nil, err
	}

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{WalletAddress: wallet.WalletAddr, Token: p.FromToken})
	if err != nil {
		return nil, err
	}

	if b.Float == 0 {
		return nil, errors.New("not enough balance in " + p.FromToken.String())
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	swap, err := swapper.MuteIOSwap(ctx, &zksyncera.MuteIOSwapReq{
		Gas:          nil,
		Amount:       defi.Percent(am, 90),
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     wallet.PK,
		EstimateOnly: true,
	})
	if err != nil {
		return nil, err
	}

	if p.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.ECost.TotalGasWei, am)
	}

	return GasStation(swap.ECost, p.Network), nil
}
