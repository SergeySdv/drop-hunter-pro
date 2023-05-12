package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/stargate/router"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type GetStargateBridgeFeeReq struct {
	ToChain v1.Network
	Wallet  common.Address
}

type GetStargateBridgeFeeRes struct {
	Fee1 *big.Int
	Fee2 *big.Int
}

func (c *EtheriumClient) GetStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error) {
	return lib.Retry[*GetStargateBridgeFeeReq, *GetStargateBridgeFeeRes](ctx, c.getStargateBridgeFee, req, &lib.RetryOpt{
		RetryCount: RetryMax,
	})
}

func (c *EtheriumClient) getStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error) {
	trx, err := router.NewRouterCaller(c.c.Dict.Stargate.StargateRouterAddress, c.cli)
	if err != nil {
		return nil, err
	}

	toAddress := req.Wallet.Bytes()
	payload := common.HexToAddress("0").Bytes()
	distChain, ok := ChainIdMap[req.ToChain]
	if !ok {
		return nil, errors.New("invalid chain: " + string(req.ToChain))
	}

	fee1, fee2, err := trx.QuoteLayerZeroFee(nil, distChain, typeFuncSwap, toAddress, payload, router.IStargateRouterlzTxObj{
		DstGasForCall:   big.NewInt(0),
		DstNativeAmount: big.NewInt(0),
		DstNativeAddr:   []byte{},
	})
	if err != nil {
		return nil, err
	}

	return &GetStargateBridgeFeeRes{
		Fee1: fee1,
		Fee2: fee2,
	}, nil
}
