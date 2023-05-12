package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/stargate/routereth"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type StargateBridgeSwapEthReq struct {
	DestChain v1.Network
	Wallet    *WalletTransactor
	Quantity  *big.Int
	Fee       *big.Int
	Gas       *GasLimit
	Opt       *lib.RetryOpt
}

func (r *StargateBridgeSwapEthReq) Validate(srcChain v1.Network) error {
	if srcChain == r.DestChain {
		return errors.New("same chain: " + string(srcChain))
	}

	if r.Wallet == nil {
		return errors.New("wallet is not set")
	}

	if r.Quantity == nil || r.Quantity.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("quantity is not set")
	}
	return nil
}

type StargateBridgeSwapEthRes struct {
	Tx *types.Transaction
}

func (c *EtheriumClient) StargateBridgeSwapEth(ctx context.Context, req *StargateBridgeSwapEthReq) (*StargateBridgeSwapEthRes, error) {

	if err := req.Validate(c.c.Network); err != nil {
		return nil, err
	}

	return lib.Retry[*StargateBridgeSwapEthReq, *StargateBridgeSwapEthRes](ctx, c.stargateBridgeSwapEth, req, req.Opt)
}
func (c *EtheriumClient) stargateBridgeSwapEth(ctx context.Context, req *StargateBridgeSwapEthReq) (*StargateBridgeSwapEthRes, error) {

	tr, err := routereth.NewRouterTransactor(c.c.Dict.Stargate.StargateRouterEthAddress, c.cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.cli.ChainID(ctx)
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

	destChainId := ChainIdMap[req.DestChain]

	prec := big.NewInt(0).Div(req.Quantity, big.NewInt(1000))
	min := big.NewInt(0).Mul(prec, big.NewInt(995))

	opt.Value = big.NewInt(0).Add(req.Quantity, req.Fee)
	opt.NoSend = true
	tx, err := tr.SwapETH(
		opt,
		destChainId,
		req.Wallet.WalletAddr,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		min,
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SwapETH")
	}

	opt.GasPrice = c.ResolveGasPrice(req.Gas, tx)
	opt.NoSend = false

	tx, err = tr.SwapETH(
		opt,
		destChainId,
		req.Wallet.WalletAddr,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		min,
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SwapETH")
	}

	return &StargateBridgeSwapEthRes{
		tx,
	}, nil
}
