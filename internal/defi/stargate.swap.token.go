package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/router"
	"github.com/pkg/errors"
)

type StargateBridgeSwapTokenReq = StargateBridgeSwapReq

type StargateBridgeSwapToken struct {
	Tx    *types.Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapToken(ctx context.Context, req *StargateBridgeSwapTokenReq) (*StargateBridgeSwapToken, error) {

	if err := req.Validate(c.Cfg.Network); err != nil {
		return nil, err
	}

	tr, err := router.NewRouterTransactor(c.Cfg.Dict.Stargate.StargateRouterAddress, c.Cli)
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

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.DestChain,
		Wallet:  req.Wallet.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	opt.Value = fee.Fee1
	destChainId := ChainIdMap[req.DestChain]
	srcPoolId := PoolIdMap[c.Cfg.Network][req.FromToken]
	distPoolId := PoolIdMap[req.DestChain][req.ToToken]
	min := Slippage(req.Quantity, SlippagePercent05)
	opt.NoSend = req.EstimateOnly

	if req.Gas.RuleSet() {
		opt.GasLimit = req.Gas.GasLimit.Uint64()
		opt.GasPrice = &req.Gas.GasPrice
	}

	// 38.677058
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
			DstNativeAddr:   common.HexToAddress("0x0000000000000000000000000000000000000001").Bytes(),
		},
		req.Wallet.WalletAddr.Bytes(),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.Swap")
	}

	eCost := &EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(tx.Gas()),
		GasPrice:    tx.GasPrice(),
		TotalGasWei: new(big.Int).Add(MinerGasLegacy(tx.GasPrice(), tx.Gas()), fee.Fee1),
	}

	return &StargateBridgeSwapToken{
		Tx:    tx,
		ECost: eCost,
	}, nil
}
