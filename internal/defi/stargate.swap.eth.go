package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/routereth"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapEthReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	Gas          *Gas
	EstimateOnly bool
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
	Tx    *types.Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapEth(ctx context.Context, req *StargateBridgeSwapEthReq) (*StargateBridgeSwapEthRes, error) {

	//var gasLimitPrice = map[v1.Network]uint64{
	//	v1.Network_ARBITRUM:    4_000_000,
	//	v1.Network_OPTIMISM:    1_000_000,
	//	v1.Network_Etherium:    630_000,
	//	v1.Network_POLIGON:     630_000,
	//	v1.Network_BinanaceBNB: 600_000,
	//	v1.Network_AVALANCHE:   700_000,
	//}

	if err := req.Validate(c.Cfg.Network); err != nil {
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

	destChainId := ChainIdMap[req.DestChain]

	l1fee := big.NewInt(0)
	if c.Cfg.Network == v1.Network_OPTIMISM {
		optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.Cli)
		if err != nil {
			return nil, err
		}

		abi, err := routereth.StorageMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		data, err := abi.Pack("swapETH",
			destChainId,
			req.Wallet.WalletAddr,
			req.Wallet.WalletAddr.Bytes(),
			req.Quantity,
			Slippage(req.Quantity, SlippagePercent05))
		if err != nil {
			return nil, err
		}
		o := &bind.CallOpts{}
		o.Context = ctx
		l1fee, err = optFeeCaller.GetL1Fee(o, data)
		if err != nil {
			return nil, err
		}

		var gasPrice *big.Int
		var optimismDefaultGasPrice int64 = 100_000_058
		maxGasPriceRetry := 3
		counter := 0
		for counter < maxGasPriceRetry {
			counter++
			gasPrice, err = c.SuggestGasPrice(ctx)
			if err != nil {
				continue
			}
			if gasPrice.Cmp(big.NewInt(optimismDefaultGasPrice)) < 0 {
				continue
			}
		}
		if gasPrice.Cmp(big.NewInt(optimismDefaultGasPrice)) == 0 {
			gasPrice = big.NewInt(optimismDefaultGasPrice)
		}

		opt.GasPrice = gasPrice
	}

	opt.NoSend = req.EstimateOnly

	if req.Gas.RuleSet() {
		opt.GasLimit = req.Gas.GasLimit.Uint64()
		opt.GasPrice = &req.Gas.GasPrice
		opt.GasPrice = ResolveGasPriceLayerZero(&req.Gas.MaxGas, &req.Gas.GasLimit, &req.Gas.GasPrice)
	}

	opt.Value = BigIntSum(req.Quantity, fee.Fee1)

	opt.NoSend = req.EstimateOnly

	tr, err := routereth.NewStorageTransactor(c.Cfg.Dict.Stargate.StargateRouterEthAddress, c.Cli)
	if err != nil {
		return nil, err
	}
	tx, err := tr.SwapETH(
		opt,
		destChainId,
		req.Wallet.WalletAddr,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		Slippage(req.Quantity, SlippagePercent05),
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SwapETH")
	}

	// am 1799274916775232
	//  b 2338339511083492449
	//  v 2200050382441954
	//  g 460521821676035
	//

	e := &EstimatedGasCost{
		GasLimit:    new(big.Int).SetUint64(tx.Gas()),
		GasPrice:    GasPrice(tx),
		TotalGasWei: BigIntSum(MinerGasLegacy(GasPrice(tx), tx.Gas()), fee.Fee1, l1fee),
	}

	return &StargateBridgeSwapEthRes{
		Tx:    tx,
		ECost: e,
	}, nil
}

func GasPrice(tx *types.Transaction) *big.Int {
	if tx.GasFeeCap() != nil || tx.GasTipCap().Cmp(big.NewInt(0)) > 0 {
		return tx.GasFeeCap()
	}
	return tx.GasPrice()
}

func BigIntSum(values ...*big.Int) *big.Int {

	result := big.NewInt(0)

	for _, v := range values {
		result = new(big.Int).Add(v, result)
	}

	return result

}

type SlippagePercent = string

const (
	SlippagePercent1   SlippagePercent = "1"
	SlippagePercent05  SlippagePercent = "0.5"
	SlippagePercent02  SlippagePercent = "0.2"
	SlippagePercent03  SlippagePercent = "0.3"
	SlippagePercent01  SlippagePercent = "0.1"
	SlippagePercent001 SlippagePercent = "0.01"
)

func Slippage(v *big.Int, slippagePercent SlippagePercent) *big.Int {
	switch slippagePercent {
	case SlippagePercent05:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(995))
	case SlippagePercent01:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(999))
	case SlippagePercent001:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9999))
	case SlippagePercent02:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9998))
	case SlippagePercent03:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9997))
	case SlippagePercent1:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(99))
	}
	return nil
}
