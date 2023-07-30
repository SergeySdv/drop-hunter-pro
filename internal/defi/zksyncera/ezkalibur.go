package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/ezkaliburrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) EzkaliburSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error) {
	result := &defi.DefaultSwapRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.Ezkalibur.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTxHash = tokenLimitChecker.ApproveTx

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	data, err := c.makeEzkaliburSwapData(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "makeEzkaliburSwapData")
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.Ezkalibur.Router,
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

func (c *Client) makeEzkaliburSwapData(ctx context.Context, req *defi.DefaultSwapReq) ([]byte, error) {

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	call, err := ezkaliburrouter.NewStorageCaller(c.Cfg.Ezkalibur.Router, c.Provider)
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

	abiEzkaliburrouter, err := ezkaliburrouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if req.FromToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent05)
		return abiEzkaliburrouter.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amOut, path, w.WalletAddr, w.WalletAddr, DefaultDeadLine())

	} else if req.ToToken == v1.Token_ETH {
		amOut := amsOut[1]
		amOut = defi.Slippage(amOut, defi.SlippagePercent05)
		return abiEzkaliburrouter.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", req.Amount, amOut, path, w.WalletAddr, w.WalletAddr, DefaultDeadLine())
	}

	return nil, errors.New("unsupported input params")
}
