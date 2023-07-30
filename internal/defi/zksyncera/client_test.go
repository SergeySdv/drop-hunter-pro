package zksyncera

import (
	"context"
	"testing"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	r, err := NewMainNetClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	ctx := context.Background()

	poolAddr, err := r.GetSyncSwapPool(ctx, &GetSyncSwapPoolReq{
		A: v1.Token_USDC,
		B: v1.Token_ETH,
	})
	assert.NoError(t, err)
	assert.NotNil(t, poolAddr)

	b, err := r.SyncSwapLPBalance(ctx, *poolAddr, tests.GetConfig().Wallet)
	assert.NoError(t, err)
	assert.NotNil(t, b)

	res, err := r.SyncSwapLiquidity(ctx, &SyncSwapLiquidityReq{
		Amount:       b,
		A:            v1.Token_USDC,
		B:            v1.Token_ETH,
		WalletPK:     tests.GetConfig().PK,
		Add:          false,
		EstimateOnly: true,
		Gas:          nil,
		debug:        true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	//
	//err = r.WaitTxComplete(, common.HexToHash("0x56d5d3590fbf244388a35c225e8389e17f30ab4c5c55b20bf1665bf120d3bdf4"))
	//assert.NoError(t, err)
	//
	//wallet, err := NewWalletTransactor(tests.GetConfig().PK, r.NetworkId)
	//assert.NoError(t, err)
	//
	//am := big.NewInt(1)
	//res, err := r.TokenLimitChecker(context.Background(), &TokenLimitCheckerReq{
	//	Token:       v1.Token_USDC,
	//	Wallet:      wallet,
	//	Amount:      am,
	//	SpenderAddr: r.Cfg.SyncSwap.RouterSwap,
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, res)
	//
	//swap, err := r.SyncSwap(context.Background(), &defi.SyncSwapReq{
	//	Network:   v1.Network_ZKSYNCERATESTNET,
	//	Amount:    am,
	//	FromToken: v1.Token_USDC,
	//	ToToken:   v1.Token_ETH,
	//	WalletPK:  tests.GetConfig().PK,
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, swap)

}
