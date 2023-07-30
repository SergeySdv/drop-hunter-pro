package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type SyncSwapTask struct {
	cancel func()
}

func (t *SyncSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *SyncSwapTask) Reset(ctx context.Context, a *Input) error {
	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_SyncSwapTask)
	if !ok {
		return errors.New("panic.a.Task.Task.Task.(*v1.Task_SyncSwapTask) call an ambulance!")
	}

	p := l.SyncSwapTask

	p.TxId = nil
	p.TxCompleted = nil
	p.SwapCompleted = nil

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *SyncSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_SyncSwapTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_SyncSwapTask) call an ambulance!")
	}

	p := l.SyncSwapTask

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

	client, _, err := NewSyncSwapper(taskContext, a)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		res, gas, err := Swap(taskContext, p, client, profile)
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(res.Swap, gas)
		if err := a.AddTx(ctx, res.Allowance); err != nil {
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

func NewSyncSwapper(ctx context.Context, a *Input) (defi.SyncSwapper, *zksyncera.WalletTransactor, error) {

	l, ok := a.Task.Task.Task.(*v1.Task_SyncSwapTask)
	if !ok {
		return nil, nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, nil, err
	}

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	stgs, err := a.SettingsService.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: l.SyncSwapTask.Network,
		UserId:  a.UserId,
	})
	if err != nil {
		return nil, nil, err
	}

	swapper, err := uniclient.NewSyncSwapper(l.SyncSwapTask.Network, &uniclient.BaseClientConfig{
		ProxyString:     proxyString,
		RPCEndpoint:     stgs.RpcEndpoint,
		UserAgentHeader: profile.UserAgent,
	})
	if err != nil {
		return nil, nil, err
	}
	wallet, err := zksyncera.NewWalletTransactor(string(profile.MmskPk), swapper.GetNetworkId())
	if err != nil {
		return nil, nil, err
	}

	return swapper, wallet, nil
}

type SyncSwapRes struct {
	Swap      *defi.Transaction
	Allowance *defi.Transaction
	Swapper   defi.SyncSwapper
}

func Swap(ctx context.Context, a *v1.SyncSwapTask, swapper defi.SyncSwapper, profile *halp.Profile) (*SyncSwapRes, *defi.Gas, error) {

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: profile.WalletAddr,
		Token:         a.FromToken,
	})
	if err != nil {
		return nil, nil, err
	}

	if b.Float == 0 {
		return nil, nil, errors.New("not enough balance in " + a.FromToken.String())
	}

	am, err := defi.ResolveAmount(a.Amount, b.WEI)
	if err != nil {
		return nil, nil, err
	}

	estimation, err := EstimateSyncSwapCost(ctx, profile, a)
	if err != nil {
		return nil, nil, err
	}

	gas, err := GasManager(estimation, profile.Settings, a.Network)
	if err != nil {
		return nil, nil, err
	}
	if a.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
	}

	swap, err := swapper.SyncSwap(ctx, &defi.SyncSwapReq{
		Network:   a.Network,
		Amount:    am,
		FromToken: a.FromToken,
		ToToken:   a.ToToken,
		WalletPK:  profile.WalletPK,
		Gas:       gas,
	})
	if err != nil {
		return nil, nil, err
	}

	result := &SyncSwapRes{
		Swap:      swap.SwapTx,
		Allowance: swap.ApproveTxHash,
		Swapper:   swapper,
	}

	return result, gas, nil
}

func EstimateSyncSwapCost(ctx context.Context, profile *halp.Profile, p *v1.SyncSwapTask) (*v1.EstimationTx, error) {

	swapper, err := uniclient.NewSyncSwapper(p.Network, profile.BaseConfig(p.Network))
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

	swap, err := swapper.SyncSwap(ctx, &defi.SyncSwapReq{
		Network:      p.Network,
		Amount:       defi.Percent(am, 50),
		FromToken:    p.FromToken,
		ToToken:      p.ToToken,
		WalletPK:     wallet.PK,
		EstimateOnly: true,
	})
	if err != nil {
		return nil, err
	}

	if p.FromToken == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.EstimatedGasCost.TotalGasWei, am)
	}

	return GasStation(swap.EstimatedGasCost, p.Network), nil
}
