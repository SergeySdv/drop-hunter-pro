package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/stargate/router"
	"crypto_scripts/internal/lib"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type StargateBridgeSwapTokenReq = StargateBridgeSwapReq

type StargateBridgeSwapToken struct {
	Tx *types.Transaction
}

func (c *EtheriumClient) StargateBridgeSwapToken(ctx context.Context, req *StargateBridgeSwapTokenReq) (*StargateBridgeSwapToken, error) {

	if err := req.Validate(c.c.Network); err != nil {
		return nil, err
	}

	return lib.Retry[*StargateBridgeSwapReq, *StargateBridgeSwapToken](ctx, c.stargateBridgeSwapOther, req, req.Opt)
}
func (c *EtheriumClient) stargateBridgeSwapOther(ctx context.Context, req *StargateBridgeSwapTokenReq) (*StargateBridgeSwapToken, error) {

	tr, err := router.NewRouterTransactor(c.c.Dict.Stargate.StargateRouterAddress, c.cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.GetNetworkId(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	if req.Fee == nil {
		fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
			ToChain: req.DestChain,
			Wallet:  req.Wallet.WalletAddr,
		})
		if err != nil {
			return nil, errors.Wrap(err, "GetStargateBridgeFee")
		}
		req.Fee = fee.Fee1
	}

	opt.Value = req.Fee

	destChainId := ChainIdMap[req.DestChain]
	srcPoolId := PoolIdMap[c.c.Network][req.FromToken]
	distPoolId := PoolIdMap[req.DestChain][req.ToToken]

	prec := big.NewInt(0).Div(req.Quantity, big.NewInt(1000))
	min := big.NewInt(0).Mul(prec, big.NewInt(995))
	opt.NoSend = true
	tx, err := tr.Swap(
		opt,
		destChainId,
		big.NewInt(srcPoolId),
		big.NewInt(distPoolId),
		req.Wallet.WalletAddr,
		req.Quantity,
		min,
		router.IStargateRouterlzTxObj{
			DstGasForCall:   big.NewInt(0),
			DstNativeAmount: big.NewInt(0),
			DstNativeAddr:   common.HexToAddress("0x").Bytes(),
		},
		req.Wallet.WalletAddr.Bytes(),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.Swap")
	}

	println("tx.GasPrice()", tx.GasPrice())
	println("tx.Gas()", tx.Gas())
	println("tx.GasFeeCap()", tx.GasFeeCap())
	println("tx.GasTipCap()", tx.GasTipCap())

	opt.GasPrice = c.ResolveGasPrice(req.Gas, tx)

	println("gas price after ", opt.GasPrice.String())
	opt.NoSend = false

	tx, err = tr.Swap(
		opt,
		destChainId,
		big.NewInt(srcPoolId),
		big.NewInt(distPoolId),
		req.Wallet.WalletAddr,
		req.Quantity,
		min,
		router.IStargateRouterlzTxObj{
			DstGasForCall:   big.NewInt(0),
			DstNativeAmount: big.NewInt(0),
			DstNativeAddr:   common.HexToAddress("0x").Bytes(),
		},
		req.Wallet.WalletAddr.Bytes(),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.Swap")
	}

	return &StargateBridgeSwapToken{
		tx,
	}, nil
}
