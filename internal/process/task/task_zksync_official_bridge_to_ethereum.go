package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type ZksyncOfficialBridgeToEthereumTask struct {
	cancel func()
}

func (t *ZksyncOfficialBridgeToEthereumTask) Stop() error {
	t.cancel()
	return nil
}

func (t *ZksyncOfficialBridgeToEthereumTask) Reset(ctx context.Context, a *Input) error {
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

func (t *ZksyncOfficialBridgeToEthereumTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask) call an ambulance!")
	}

	p := l.ZkSyncOfficialBridgeToEthereumTask

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning:
		task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	client, wallet, err := NewZkSyncClient(profile, p.Network)
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {
		tx, gas, err := t.Withdrawal(taskContext, a, client, wallet, profile)
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(tx, gas)
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

func NewZkSyncClient(profile *halp.Profile, n v1.Network) (*zksyncera.Client, *zksyncera.WalletTransactor, error) {
	swapper, err := uniclient.NewZkSyncOfficialBridge(n, profile.BaseConfig(n))
	if err != nil {
		return nil, nil, err
	}

	wallet, err := zksyncera.NewWalletTransactor(profile.WalletPK, swapper.GetNetworkId())
	if err != nil {
		return nil, nil, err
	}

	return swapper, wallet, nil
}

func (t *ZksyncOfficialBridgeToEthereumTask) Withdrawal(ctx context.Context, a *Input, client *zksyncera.Client, wallet *zksyncera.WalletTransactor, profile *halp.Profile) (*defi.Transaction, *defi.Gas, error) {

	l, ok := a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask)
	if !ok {
		return nil, nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_ZkSyncOfficialBridgeToEthereumTask) call an ambulance!")
	}

	b, err := client.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: wallet.WalletAddr,
		Token:         v1.Token_ETH,
	})
	if err != nil {
		return nil, nil, err
	}

	if b.Float == 0 {
		return nil, nil, errors.New("not enough balance in ETH")
	}

	am, err := defi.ResolveAmount(l.ZkSyncOfficialBridgeToEthereumTask.Amount, b.WEI)
	if err != nil {
		return nil, nil, err
	}

	estimate, err := client.BridgeToEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:       defi.Percent(am, 50),
		WalletPK:     wallet.PK,
		EstimateOnly: true,
	})
	if err != nil {
		return nil, nil, err
	}

	gasStation := GasStation(estimate.EstimatedGasCost, v1.Network_ZKSYNCERA)
	gas, err := GasManager(gasStation, profile.Settings, v1.Network_ZKSYNCERA)
	if err != nil {
		return nil, nil, err
	}

	am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)

	swap, err := client.BridgeToEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:       am,
		WalletPK:     wallet.PK,
		Gas:          gas,
		EstimateOnly: false,
	})
	if err != nil {
		return nil, nil, err
	}

	result := swap.TxHash

	return result, gas, nil
}

func EstimateZkSyncOfficialBridgeToEthSwapCost(ctx context.Context, profile *halp.Profile, l *v1.ZkSyncOfficialBridgeToEthereumTask) (*v1.EstimationTx, error) {

	network := l.Network

	client, err := uniclient.NewZkSyncOfficialBridge(network, profile.BaseConfig(network))
	if err != nil {
		return nil, err
	}

	b, err := client.GetBalance(ctx, &defi.GetBalanceReq{WalletAddress: profile.WalletAddr, Token: v1.Token_ETH})
	if err != nil {
		return nil, err
	}

	am, err := defi.ResolveAmount(l.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	res, err := client.BridgeToEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:       defi.Percent(am, 50),
		WalletPK:     profile.WalletPK,
		EstimateOnly: true,
	})
	if err != nil {
		return nil, err
	}

	am = ResolveNetworkTokenAmount(b.WEI, res.EstimatedGasCost.TotalGasWei, am)

	res, err = client.BridgeToEthereumNetwork(ctx, &zksyncera.L1L2BridgeReq{
		Amount:       am,
		WalletPK:     profile.WalletPK,
		EstimateOnly: true,
	})
	if err != nil {
		return nil, err
	}

	return &v1.EstimationTx{
		Balance:  defi.AmountUni(b.WEI, network),
		Value:    defi.AmountUni(am, network),
		Gas:      defi.AmountUni(res.EstimatedGasCost.TotalGasWei, network),
		GasLimit: defi.AmountUni(res.EstimatedGasCost.GasLimit, network),
		GasPrice: defi.AmountUni(res.EstimatedGasCost.GasPrice, network),
	}, nil
}
