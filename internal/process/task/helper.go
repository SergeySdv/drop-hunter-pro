package task

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/pkg/errors"
)

func GasStation(ecost *defi.EstimatedGasCost, network v1.Network) *v1.EstimationTx {
	result := &v1.EstimationTx{}
	result.GasLimit = defi.AmountUni(ecost.GasLimit, network)
	result.GasPrice = defi.AmountUni(ecost.GasPrice, network)
	result.Gas = defi.AmountUni(ecost.TotalGasWei, network)

	result.GasValuePercent = ""

	return result
}

func ResolveNetworkTokenAmount(balance, gas, value *big.Int) *big.Int {
	need := new(big.Int).Add(gas, value)
	if need.Cmp(balance) <= 0 {
		return value
	}

	return defi.Percent(new(big.Int).Sub(balance, gas), 99)
}

func WaitTxComplete(ctx context.Context, ptx *v1.TaskTx, task *v1.ProcessTask, networker defi.Networker, updater TaskUpdater) error {
	if ptx != nil && ptx.TxId != "" {
		tx := ptx
		if err := networker.WaitTxComplete(ctx, common.HexToHash(tx.TxId)); err != nil {
			if errors.Is(err, defi.ErrTxStatusFailed) {
				tx.TxCompleted = false
				tx.RetryCount++
				if tx.RetryCount > 2 {
					tx.TxCompleted = false
					tx.RetryCount = 0
					tx.TxId = ""
				}
				if err := updater.UpdateTask(ctx, task); err != nil {
					return err
				}
			}
			if errors.Is(err, defi.ErrTxNotFound) {
				if tx.RetryCount > 2 {
					tx.TxCompleted = false
					tx.RetryCount = 0
					tx.TxId = ""
				}
				tx.RetryCount++
				if err := updater.UpdateTask(ctx, task); err != nil {
					return err
				}
			}
			return err
		}

		tx.TxCompleted = true
		if err := updater.UpdateTask(ctx, task); err != nil {
			return err
		}
	}
	return nil
}

func GasMultiplier(mul *float64, gas *defi.Gas) *defi.Gas {

	if mul == nil || *mul == 0 || *mul == 1 {
		gas.GasMultiplier = 1
		return gas
	} else {
		gas.GasMultiplier = *mul
	}

	f := new(big.Float).SetFloat64(*mul)

	newGasLimit := new(big.Float).Mul(big.NewFloat(0).SetInt64(gas.GasLimit.Int64()), f)
	gl, _ := newGasLimit.Int(nil)
	gas.GasLimit = *gl

	total := big.NewInt(0).Mul(&gas.GasLimit, &gas.GasPrice)
	gas.TotalGas = *total

	return gas
}

func (a *Input) AddTx2(ctx context.Context, transactions ...*v1.TaskTx) error {

	for i := range transactions {
		tx := transactions[i]
		if tx == nil || tx.TxCompleted == false {
			continue
		}
		txDb := &repository.Transaction{
			Id:        tx.GetTxId(),
			Code:      tx.GetCode(),
			Network:   tx.GetNetwork().String(),
			Url:       tx.GetUrl(),
			TaskId:    a.Task.Id,
			ProfileId: a.ProfileId,
			ProcessId: a.ProcessId,
			UserId:    a.UserId,
			CreatedAt: time.Now(),
		}
		if err := TransactionAdd(ctx, a.ProcessRepository, txDb); err != nil {
			return err
		}
	}

	return nil
}

func (a *Input) AddTx(ctx context.Context, transactions ...*defi.Transaction) error {

	for i := range transactions {
		tx := transactions[i]
		if tx == nil {
			continue
		}
		txDb := &repository.Transaction{
			Id:        tx.Id,
			Code:      tx.Code,
			Network:   tx.Network.String(),
			Url:       tx.Url,
			TaskId:    a.Task.Id,
			ProfileId: a.ProfileId,
			ProcessId: a.ProcessId,
			UserId:    a.UserId,
			CreatedAt: time.Now(),
		}
		if err := TransactionAdd(ctx, a.ProcessRepository, txDb); err != nil {
			return err
		}
	}

	return nil
}

func TransactionAdd(ctx context.Context, rep repository.TransactionRepository, tx *repository.Transaction) error {
	return rep.TransactionAdd(ctx, tx)
}

func NewTx(tx *defi.Transaction, gas *defi.Gas) *v1.TaskTx {

	if tx == nil {
		return nil
	}
	txx := &v1.TaskTx{
		TxCompleted: false,
		TxId:        tx.Id,
		RetryCount:  0,
		Url:         &tx.Url,
		Network:     &tx.Network,
		Code:        &tx.Code,
	}

	if gas != nil {
		txx.GasEstimated = defi.AmountUni(&gas.GasBeforeMultiplier, gas.Network)
		txx.GasResult = defi.AmountUni(&gas.TotalGas, gas.Network)
		txx.GasLimit = defi.AmountUni(&gas.MaxGas, gas.Network)
		m := float32(gas.GasMultiplier)
		txx.Multiplier = &m
	}
	return txx
}
