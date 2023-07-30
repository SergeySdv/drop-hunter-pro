package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/contracts/testnetbridge"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type TestNetBridgeSwapReq struct {
	Network Network
	Wallet  *WalletTransactor
	Amount  *big.Int
	Gas     *GasLimit
}

type TestNetBridgeSwapRes struct {
	Tx *Transaction
}

var testNetBridgeSwapDist = map[Network]uint16{
	v1.Network_ARBITRUM: 154,
	v1.Network_Etherium: 154,
	v1.Network_OPTIMISM: 154,
}

func (r *TestNetBridgeSwapReq) Validate() error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Wallet == nil {
		return errors.New("wallet is empty")
	}

	if r.Amount.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("Amount is zero ir negative")
	}

	_, exist := testNetBridgeSwapDist[r.Network]
	if !exist {
		return errors.New("network is not supported: " + r.Network.String())
	}

	return nil
}

// ARBITRUM
// {
//  "name": "swapAndBridge",
//  "params": [
//    {
//      "name": "amountIn",
//      "value": "1000000000000000",
//      "type": "uint256"
//    },
//    {
//      "name": "amountOutMin",
//      "value": "14506105063062191234",
//      "type": "uint256"
//    },
//    {
//      "name": "dstChainId",
//      "value": "154",
//      "type": "uint16"
//    },
//    {
//      "name": "to",
//      "value": "0x4a6e7c137a6691d55693ca3bc7e5c698d9d43815",
//      "type": "address"
//    },
//    {
//      "name": "refundAddress",
//      "value": "0x4a6e7c137a6691d55693ca3bc7e5c698d9d43815",
//      "type": "address"
//    },
//    {
//      "name": "zroPaymentAddress",
//      "value": "0x0000000000000000000000000000000000000000",
//      "type": "address"
//    },
//    {
//      "name": "adapterParams",
//      "value": null,
//      "type": "bytes"
//    }
//  ]
//}

func (c *EtheriumClient) TestNetBridgeSwap(ctx context.Context, req *TestNetBridgeSwapReq) (*TestNetBridgeSwapRes, error) {

	if err := req.Validate(); err != nil {
		return nil, errors.Wrap(err, "TestNetBridgeSwap")
	}

	trx, err := testnetbridge.NewTestnetbridgeTransactor(c.Cfg.Dict.TestNetBridgeSwapAddress, c.Cli)
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

	goerliEthAmount := big.NewInt(0).Mul(req.Amount, big.NewInt(5000))
	dstChainId := testNetBridgeSwapDist[req.Network]
	zeroPaymentAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: v1.Network_AVALANCHE,
		Wallet:  req.Wallet.WalletAddr,
	})
	if err != nil {
		return nil, err
	}

	opt.Value = big.NewInt(0).Add(req.Amount, fee.Fee1)

	tx, err := trx.SwapAndBridge(opt, req.Amount, goerliEthAmount, dstChainId, req.Wallet.WalletAddr, req.Wallet.WalletAddr, zeroPaymentAddr, nil)
	if err != nil {
		return nil, err
	}

	return &TestNetBridgeSwapRes{
		Tx: c.NewTx(tx.Hash(), CodeContract),
	}, nil
}
