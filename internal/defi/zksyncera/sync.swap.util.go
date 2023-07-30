package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/syncswapclassicpool"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/syncswappoolfactory"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type getSyncSwapFeeReq struct {
	Network   v1.Network
	FromToken v1.Token
	ToToken   v1.Token
	Wallet    common.Address
}

type getSyncSwapFeeRes struct {
	Fee      *big.Int
	PoolAddr common.Address
}

func (c *Client) getSyncSwapFee(ctx context.Context, req *getSyncSwapFeeReq) (*getSyncSwapFeeRes, error) {
	cstorage, err := syncswappoolfactory.NewStorageCaller(c.Cfg.SyncSwap.ClassicPoolFactory, c.Provider)
	if err != nil {
		return nil, err
	}

	fromToken := c.Cfg.TokenMap[req.FromToken]
	toToken := c.Cfg.TokenMap[req.ToToken]

	opt := &bind.CallOpts{Context: ctx}
	poolAddr, err := cstorage.GetPool(opt, toToken, fromToken)
	if err != nil {
		return nil, err
	}
	if poolAddr.String() == "0x0000000000000000000000000000000000000000" {
		return nil, errors.New("pool is not exist")
	}

	sender := req.Wallet
	fee, err := cstorage.GetSwapFee(opt, poolAddr, sender, fromToken, toToken, []byte("0x"))
	if err != nil {
		return nil, err
	}

	return &getSyncSwapFeeRes{Fee: fee, PoolAddr: poolAddr}, nil
}

type getAmountOutReq struct {
	Addr      common.Address
	FromToken defi.Token
	ToToken   defi.Token
	Amount    *big.Int
}

func (c *Client) getAmountOut(ctx context.Context, req *getAmountOutReq) (*big.Int, error) {

	fromToken, supported := c.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.FromToken)
	}

	poolAddr, err := c.GetSyncSwapPool(ctx, &GetSyncSwapPoolReq{
		A: req.FromToken,
		B: req.ToToken,
	})
	if err != nil {
		return nil, err
	}
	opt := &bind.CallOpts{Context: ctx}
	trx, err := syncswapclassicpool.NewStorageCaller(*poolAddr, c.Provider)
	if err != nil {
		return nil, err
	}

	return trx.GetAmountOut(opt, fromToken, req.Amount, req.Addr)
}

func (c *Client) syncSwapPoolRates(ctx context.Context, addr common.Address) (*big.Float, *big.Float, error) {
	opt := &bind.CallOpts{Context: ctx}
	trx, err := syncswapclassicpool.NewStorageCaller(addr, c.Provider)
	if err != nil {
		return nil, nil, err
	}
	r0, err := trx.Reserve0(opt)
	if err != nil {
		return nil, nil, err
	}
	r1, err := trx.Reserve1(opt)
	if err != nil {
		return nil, nil, err
	}

	totalSuply, err := trx.TotalSupply(opt)
	if err != nil {
		return nil, nil, err
	}

	ratio0 := new(big.Float).Quo(new(big.Float).SetInt(totalSuply), new(big.Float).SetInt(r0))
	ratio2 := new(big.Float).Quo(new(big.Float).SetInt(totalSuply), new(big.Float).SetInt(r1))
	return ratio0, ratio2, nil
}

type GetSyncSwapPoolReq struct {
	A v1.Token
	B v1.Token
}

func (c *Client) GetSyncSwapPool(ctx context.Context, req *GetSyncSwapPoolReq) (*common.Address, error) {

	A, supported := c.Cfg.TokenMap[req.A]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.A)
	}
	B, supported := c.Cfg.TokenMap[req.B]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.B)
	}

	cstorage, err := syncswappoolfactory.NewStorageCaller(c.Cfg.SyncSwap.ClassicPoolFactory, c.Provider)
	if err != nil {
		return nil, err
	}
	opt := &bind.CallOpts{Context: ctx}

	addr, err := cstorage.GetPool(opt, A, B)
	if err != nil {
		return nil, err
	}
	return &addr, nil

}

func (c *Client) SyncSwapLPBalance(ctx context.Context, pool, user common.Address) (*big.Int, error) {
	trx, err := syncswapclassicpool.NewStorageCaller(pool, c.Provider)
	if err != nil {
		return nil, err
	}

	return trx.BalanceOf(&bind.CallOpts{Context: ctx}, user)
}
