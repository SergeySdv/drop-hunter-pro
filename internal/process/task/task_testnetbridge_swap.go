package task

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TestNetBridgeSwapTask struct {
	cancel func()
}

func (t *TestNetBridgeSwapTask) Stop() error {
	t.cancel()
	return nil
}

func (t *TestNetBridgeSwapTask) Reset(ctx context.Context, a *Input) error {
	return nil
}

func (t *TestNetBridgeSwapTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_TestNetBridgeSwapTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_TestNetBridgeSwapTask) call an ambulance!")
	}

	p := l.TestNetBridgeSwapTask

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning:
		a.Task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	minAmount, err := lib.StringToFloat(p.MinAmount)

	maxAmount, err := lib.StringToFloat(p.MaxAmount)

	amountFloat := lib.RandFloatRange(minAmount, maxAmount)
	amountString := lib.FloatToString(amountFloat)
	p.Amount = &amountString
	amountWei := defi.CastFloatToEtherWEI(amountFloat)

	if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
		return nil, err
	}

	wallet, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	stgs, err := a.SettingsService.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: p.Network,
		UserId:  a.UserId,
	})
	if err != nil {
		return nil, err
	}

	client, err := uniclient.NewTestNetworkBridgeSwapper(p.Network, &uniclient.BaseClientConfig{
		ProxyString:     proxyString,
		RPCEndpoint:     stgs.RpcEndpoint,
		UserAgentHeader: profile.UserAgent,
	})
	if err != nil {
		return nil, err
	}

	if p.GetTx().GetTxId() == "" {

		res, err := client.TestNetBridgeSwap(taskContext, &defi.TestNetBridgeSwapReq{
			Network: p.Network,
			Wallet:  wallet,
			Amount:  amountWei,
		})
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(res.Tx, nil)
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
		task.FinishedAt = timestamppb.Now()
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}
