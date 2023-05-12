package process

import (
	"context"
	"crypto_scripts/internal/defi"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/settings"
	"crypto_scripts/internal/uniclient"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StargateTask struct {
}

func (t *StargateTask) Run(ctx context.Context, a *TaskArg) (*TaskRes, error) {

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_StargateBridgeTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}

	p := l.StargateBridgeTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return &TaskRes{Task: a.Task}, nil
	case v1.ProcessStatus_StatusRunning:
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:

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

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	stgs, err := settings.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: l.StargateBridgeTask.FromNetwork,
		UserId:  a.UserId,
		Rep:     a.Rep,
	})
	if err != nil {
		return nil, err
	}

	swapper, err := uniclient.NewStargateSwapper(l.StargateBridgeTask.FromNetwork, &uniclient.BaseClientConfig{
		ProxyString: proxyString,
		RPCEndpoint: stgs.RpcEndpoint,
		GasLimit:    settings.GetGasLimit(stgs),
	})
	if err != nil {
		return nil, err
	}
	wallet, err := defi.NewWalletTransactor(profile.MmskPk)
	if err != nil {
		return nil, err
	}

	fromToken := l.StargateBridgeTask.FromToken

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: wallet.WalletAddr,
		Token:         fromToken,
	})
	if err != nil {
		return nil, err
	}

	if b.Float == 0 {
		return nil, errors.New("not enough balance in " + fromToken.String())
	}

	toToken := l.StargateBridgeTask.ToToken

	distChain := l.StargateBridgeTask.ToNetwork

	fee, err := swapper.GetStargateBridgeFee(ctx, &defi.GetStargateBridgeFeeReq{
		ToChain: distChain,
		Wallet:  wallet.WalletAddr,
	})
	if err != nil {
		return nil, err
	}
	feeString := fee.Fee1.String()
	p.Fee = &feeString

	swap, err := swapper.StargateBridgeSwap(ctx, &defi.StargateBridgeSwapReq{
		DestChain: distChain,
		Wallet:    wallet,
		Quantity:  b.WEI,
		FromToken: fromToken,
		ToToken:   toToken,
		Fee:       fee.Fee1,
	})
	if err != nil {
		return nil, err
	}

	task.Status = v1.ProcessStatus_StatusDone
	task.FinishedAt = timestamppb.Now()
	for _, tx := range swap.Tx {
		task.Transactions = append(task.Transactions, swapper.TxViewFn(tx.Hash().String()))
	}

	if err := a.Rep.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
		return nil, err
	}

	_ = a.Rep.RecordStatusChanged(ctx, task.Id, v1.ProcessStatus_StatusRunning, v1.ProcessStatus_StatusDone)

	return &TaskRes{
		Task: task,
	}, nil
}
