package optimism

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	err = r.WaitTxComplete(context.Background(), common.HexToHash("0x7cab0567babe1a02a0b1beb7aeb3cf6f601ad1e54cacab92f6429d7f44b48741"))
	h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, h)

	//t.Run("balance", func(t *testing.T) {
	//	t.Skip()
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         defi.TokenSTG,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})
	//
	//t.Run("stg stake", func(t *testing.T) {
	//	t.Skip()
	//	wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
	//	assert.NoError(t, err)
	//
	//	am := defi.TokenAmountToWEI(1, defi.TokenSTG)
	//
	//	res, err := r.StakeSTG(context.Background(), &defi.StakeSTGReq{
	//		Amount: am,
	//		Wallet: wallet,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, res)
	//})
	//
	//t.Run("swapFee", func(t *testing.T) {
	//	ctx := context.Background()
	//
	//	wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
	//	assert.NoError(t, err)
	//
	//	tx, err := r.StargateBridgeSwap(ctx, &defi.StargateBridgeSwapReq{
	//		DestChain: defi.SBNameArbitrum,
	//		Wallet:    wallet,
	//		Quantity:  defi.TokenAmountFloatToWEI(2, defi.TokenSTG),
	//		Token:     defi.TokenSTG,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, tx)

	//})
}
