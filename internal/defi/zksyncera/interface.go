package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
)

type LP interface {
	GetSyncSwapPool(ctx context.Context, req *GetSyncSwapPoolReq) (*common.Address, error)
	SyncSwapLiquidity(ctx context.Context, req *SyncSwapLiquidityReq) (*SyncSwapLiquidityRes, error)
	SyncSwapLPBalance(ctx context.Context, pool, user common.Address) (*big.Int, error)
	defi.Networker
}

type MaverickSwapper interface {
	MaverickSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}

type SpaceFiSwapper interface {
	SpaceFiSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}
type VelocoreSwapper interface {
	VelocoreSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}

type IzumiSwapper interface {
	IzumiSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}

type VeSyncSwapper interface {
	VeSyncSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}

type EzkaliburSwapper interface {
	EzkaliburSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}

type ZkSwapper interface {
	ZkSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error)
	defi.Networker
}
