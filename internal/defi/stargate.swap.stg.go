package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/stg"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapSTGReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	Gas          *Gas
	EstimateOnly bool
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
	Tx    *types.Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapSTG(ctx context.Context, req *StargateBridgeSwapSTGReq) (*StargateBridgeSwapSTGRes, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	tr, err := stg.NewStgTransactor(c.Cfg.TokenMap[v1.Token_STG], c.Cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.DestChain,
		Wallet:  req.Wallet.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	opt.Value = fee.Fee1
	opt.NoSend = req.EstimateOnly

	if req.Gas.RuleSet() {
		opt.GasLimit = req.Gas.GasLimit.Uint64()
		opt.GasPrice = &req.Gas.GasPrice
	}

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

	eCost := &EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(tx.Gas()),
		GasPrice:    tx.GasPrice(),
		TotalGasWei: new(big.Int).Add(MinerGasLegacy(tx.GasPrice(), tx.Gas()), fee.Fee1),
	}

	return &StargateBridgeSwapSTGRes{
		Tx:    tx,
		ECost: eCost,
	}, nil
}

func GetTxGasPrice(tx *types.Transaction) *big.Int {
	return tx.GasPrice()
}
