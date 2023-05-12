package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/stg"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type StargateBridgeSwapSTGReq struct {
	DestChain v1.Network
	Wallet    *WalletTransactor
	Quantity  *big.Int
	Fee       *big.Int
	Gas       *GasLimit
	Opt       *lib.RetryOpt
}

func (r *StargateBridgeSwapSTGReq) Validate() error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	if r.Quantity.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("zero quantity")
	}

	if r.Quantity == nil {
		return errors.New("quantity or fee empty")
	}

	return nil
}

type StargateBridgeSwapSTGRes struct {
	Tx *types.Transaction
}

func (c *EtheriumClient) StargateBridgeSwapSTG(ctx context.Context, req *StargateBridgeSwapSTGReq) (*StargateBridgeSwapSTGRes, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return lib.Retry[*StargateBridgeSwapSTGReq, *StargateBridgeSwapSTGRes](ctx, c.stargateBridgeSwapSTG, req, req.Opt)
}
func (c *EtheriumClient) stargateBridgeSwapSTG(ctx context.Context, req *StargateBridgeSwapSTGReq) (*StargateBridgeSwapSTGRes, error) {

	tr, err := stg.NewStgTransactor(c.c.TokenMap[v1.Token_STG], c.cli)
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

	opt.Value = req.Fee
	opt.NoSend = true

	destChainId := ChainIdMap[req.DestChain]

	tx, err := tr.SendTokens(
		opt,
		destChainId,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SendTokens")
	}

	opt.GasPrice = c.ResolveGasPrice(req.Gas, tx)
	opt.NoSend = false

	tx, err = tr.SendTokens(
		opt,
		destChainId,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SendTokens")
	}

	return &StargateBridgeSwapSTGRes{
		tx,
	}, nil
}
