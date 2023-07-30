package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/weth"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) SwapWETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {
	if req.Wrap {
		return c.WrapETH(ctx, req)
	} else {
		return c.UnWrapETH(ctx, req)
	}
}

func (c *Client) WrapETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {

	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	wethabi, err := weth.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("deposit")
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		c.Cfg.Weth,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &defi.WETHRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}
func (c *Client) UnWrapETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {

	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	wethabi, err := weth.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("withdraw", req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		c.Cfg.Weth,
		big.NewInt(0),
		big.NewInt(0),
		nil,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &defi.WETHRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

func (c *Client) Make712Tx(ctx context.Context, tx *types.Transaction, gasOpt *defi.Gas, signer *accounts.DefaultEthSigner) ([]byte, *defi.EstimatedGasCost, error) {
	nonce, err := c.Provider.NonceAt(ctx, tx.From, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	var gas, gasPrice *big.Int
	if gasOpt.RuleSet() {
		gas = &gasOpt.GasLimit
		gasPrice = &gasOpt.GasPrice
		gasPrice = defi.ResolveGasPriceZksync(&gasOpt.MaxGas, gas, gasPrice)
	} else {
		gasPrice, err = c.Provider.GetGasPrice()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to GetGasPrice: %w", err)
		}
		tx.GasPrice = (*hexutil.Big)(gasPrice)
		gas, err = c.Provider.EstimateGas712(tx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to EstimateGas: %w", err)
		}
	}

	prepared := &types.Transaction712{
		Nonce:      new(big.Int).SetUint64(nonce),
		GasTipCap:  big.NewInt(100_000_000), // TODO: Estimate correct one
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

	signature, err := signer.SignTypedData(signer.GetDomain(), prepared)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Signer.SignTypedData")
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return nil, nil, errors.Wrap(err, "prepared.RLPValues")
	}

	return rawTx, &defi.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}, nil

}
