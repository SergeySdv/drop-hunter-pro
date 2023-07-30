package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/syncswaprouter"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/scan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) WaitSwapFinish(ctx context.Context, txId common.Hash) (*scan.TxStatus, error) {
	return c.scanner.WaitComplete(ctx, txId.String(), time.Second*10)
}

func (c *Client) SyncSwap(ctx context.Context, req *defi.SyncSwapReq) (*defi.SyncSwapRes, error) {

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	res, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.SyncSwap.RouterSwap,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}

	fee, err := c.getSyncSwapFee(ctx, &getSyncSwapFeeReq{
		Network:   req.Network,
		FromToken: req.FromToken,
		ToToken:   req.ToToken,
		Wallet:    transactor.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetSyncSwapFee")
	}

	amOut, err := c.getAmountOut(ctx, &getAmountOutReq{
		Addr:      transactor.WalletAddr,
		FromToken: req.FromToken,
		ToToken:   req.ToToken,
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, errors.Wrap(err, "getAmountOut")
	}

	min := defi.Slippage(amOut, defi.SlippagePercent01)

	structThing, err := abi.NewType("tuple", "struct thing", []abi.ArgumentMarshaling{
		{Name: "a", Type: "address"},
		{Name: "b", Type: "address"},
		{Name: "c", Type: "uint8"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "abi.NewType")
	}

	record := struct {
		A common.Address
		B common.Address
		C uint8
	}{
		A: c.Cfg.TokenMap[req.FromToken],
		B: transactor.WalletAddr,
		C: 1,
	}

	args := abi.Arguments{
		{Type: structThing, Name: "param_one"},
	}
	data, err := args.Pack(&record)
	if err != nil {
		return nil, errors.Wrap(err, "args.Pack(&record)")
	}

	TokenIn := c.Cfg.TokenMap[req.FromToken]

	value := fee.Fee

	if req.FromToken == v1.Token_ETH {
		TokenIn = common.HexToAddress("0x0000000000000000000000000000000000000000")
		value = new(big.Int).Add(fee.Fee, req.Amount)
	}

	path := []syncswaprouter.IRouterSwapPath{
		{
			Steps: []syncswaprouter.IRouterSwapStep{
				{
					Pool:         fee.PoolAddr,                                                      // +
					Data:         data,                                                              //+
					Callback:     common.HexToAddress("0x0000000000000000000000000000000000000000"), //+
					CallbackData: []byte("0x"),                                                      //+
				},
			},
			TokenIn:  TokenIn,    //+
			AmountIn: req.Amount, //+
		},
	}
	ded := time.Now().Add(time.Minute).UnixMilli()

	syncswaprouterabi, err := syncswaprouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err = syncswaprouterabi.Pack("swap", path, min, big.NewInt(ded))
	if err != nil {
		return nil, err
	}

	w, err := accounts.NewWallet(transactor.Signer, c.Provider)
	if err != nil {
		return nil, err
	}

	contract := c.Cfg.SyncSwap.RouterSwap

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		contract,
		big.NewInt(0),
		big.NewInt(0),
		value,
		data,
		nil, nil,
	)

	var gas, gasPrice *big.Int
	if req.Gas.RuleSet() {
		gas = &req.Gas.GasLimit
		gasPrice = &req.Gas.GasPrice
		gasPrice = defi.ResolveGasPriceZksync(&req.Gas.MaxGas, gas, gasPrice)
	} else {
		gas, err = c.Provider.EstimateGas712(tx)
		if err != nil {
			return nil, fmt.Errorf("failed to EstimateGas: %w", err)
		}
		gasPrice, err = c.Provider.GetGasPrice()
		if err != nil {
			return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
		}
	}

	nonce, err := c.Provider.NonceAt(ctx, w.GetAddress(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	prepared := &types.Transaction712{
		Nonce:      big.NewInt(0).SetUint64(nonce),
		GasTipCap:  big.NewInt(100_000_000),
		GasFeeCap:  gasPrice,
		Gas:        gas,
		To:         &tx.To,
		Value:      tx.Value.ToInt(),
		Data:       tx.Data,
		AccessList: nil,
		ChainID:    c.NetworkId,
		From:       &tx.From,
		Meta:       tx.Eip712Meta,
	}

	signature, err := transactor.Signer.SignTypedData(transactor.Signer.GetDomain(), prepared)
	if err != nil {
		return nil, errors.Wrap(err, "Signer.SignTypedData")
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return nil, errors.Wrap(err, "prepared.RLPValues")
	}

	result := &defi.SyncSwapRes{}
	result.EstimatedGasCost = &defi.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}
	result.ApproveTxHash = res.ApproveTx

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.SwapTx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}
