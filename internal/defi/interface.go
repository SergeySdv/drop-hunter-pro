package defi

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Token = v1.Token

type Network = v1.Network

type StargateStaker interface {
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	StakeSTG(ctx context.Context, req *StakeSTGReq) (*StakeSTGRes, error)
}

type Networker interface {
	TxViewFn(id string) string
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	GetNetworkToken() Token
	GetNetworkId() *big.Int
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	WaitTxComplete(ctx context.Context, tx common.Hash) error
}

type StargateSwapper interface {
	Networker
	StargateBridgeSwap(ctx context.Context, req *StargateBridgeSwapReq) (*StargateBridgeSwapRes, error)
	GetStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error)
}

type Transfer interface {
	TxViewFn(id string) string
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	GetNetworkToken() Token
	Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error)
}
