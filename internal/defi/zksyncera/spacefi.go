package zksyncera

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/spacefirouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

var DefaultDeadLine = func() *big.Int {
	return new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
}

// site ETH -> USDC https://explorer.zksync.io/tx/0xc9d9c33c9bdf7bcbafa45c313eb8b70c1d37f9797d06ce8e20f1fbf046a79ed5
// me ETH -> USDC https://explorer.zksync.io/tx/0xaddad23305e85e9c806621d274ce6ebbeeb4dfb633b415ea0a21cdb27520142b

// site USDC -> ETH https://explorer.zksync.io/tx/0x701ffd2a41613c5ecc6a1cb10ea19b742404a6b0bf3d88db82216d51eebf1b88
// me USDC -> ETH https://explorer.zksync.io/tx/0xee4a8f41ae9f32e69cb3205e88d2493ebedd9deac86b6cea9c247d896dafe590
func (c *Client) SpaceFiSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error) {
	result := &defi.DefaultSwapRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.SpaceFI.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTxHash = tokenLimitChecker.ApproveTx

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	data, err := c.makeSpaceFiSwapData(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.SpaceFI.Router,
		big.NewInt(0),
		big.NewInt(0),
		value,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	result.EstimatedGasCost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.SwapTx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

func (c *Client) makeSpaceFiSwapData(ctx context.Context, req *defi.DefaultSwapReq) ([]byte, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := spacefirouter.NewStorageCaller(c.Cfg.SpaceFI.Router, c.Provider)
	if err != nil {
		return nil, err
	}
	from, supported := c.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.FromToken)
	}
	to, supported := c.Cfg.TokenMap[req.ToToken]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.FromToken)
	}

	path := []common.Address{from, to}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abispacefirouter, err := spacefirouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if req.FromToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent001)
		return abispacefirouter.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent01)
		return abispacefirouter.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
	}

	return nil, errors.New("unsupported input params")
}
