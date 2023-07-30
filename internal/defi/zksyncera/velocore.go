package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/velocorerouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

// site ETH -> USDC
// me ETH -> USDC https://explorer.zksync.io/tx/0xd93d5099e7f0193778351365fae505e713e89acde63e9a0e229b27bf3c13d8b7

// site USDC -> ETH
// me USDC -> ETH https://explorer.zksync.io/tx/0xaa2dcfa8280071efa6fa4c7264902bbd5d589d2a9ab3b5b6e313cf903a0e9e18
func (c *Client) VelocoreSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error) {
	result := &defi.DefaultSwapRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.Velocore.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTxHash = tokenLimitChecker.ApproveTx

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	data, err := c.makeVelocoreSwapData(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeSpaceFiSwapData")
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.Velocore.Router,
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

func (c *Client) makeVelocoreSwapData(ctx context.Context, req *defi.DefaultSwapReq) ([]byte, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := velocorerouter.NewStorageCaller(c.Cfg.Velocore.Router, c.Provider)
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

	path := []velocorerouter.Routerroute{{
		From:   from,
		To:     to,
		Stable: false,
	}}
	amsOut, err := call.GetAmountsOut(&bind.CallOpts{Context: ctx}, req.Amount, path)
	if err != nil {
		return nil, err
	}
	if len(amsOut) != 2 {
		return nil, errors.New("errors getting rate")
	}

	abiSource, err := velocorerouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if req.FromToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent05)
		return abiSource.Pack("swapExactETHForTokens", amOut, path, w.WalletAddr, DefaultDeadLine())

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent05)
		return abiSource.Pack("swapExactTokensForETH", req.Amount, amOut, path, w.WalletAddr, DefaultDeadLine())
	}

	return nil, errors.New("unsupported input params")
}
