package arbitrum

import (
	"context"
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/tests"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	t.Run("balance", func(t *testing.T) {
		t.Skip()
		b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
			WalletAddress: tests.GetConfig().Wallet,
			Token:         defi.TokenETH,
		})
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	t.Run("stg stake", func(t *testing.T) {
		t.Skip()
		wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
		assert.NoError(t, err)

		am := defi.TokenAmountToWEI(6, defi.TokenSTG)

		res, err := r.StakeSTG(context.Background(), &defi.StakeSTGReq{
			Amount: am,
			Wallet: wallet,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("swapFee", func(t *testing.T) {
		ctx := context.Background()
		t.Skip()
		wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
		assert.NoError(t, err)

		tx, err := r.StargateBridgeSwap(ctx, &defi.StargateBridgeSwapReq{
			DestChain: defi.SBNameOptimism,
			Wallet:    wallet,
			Quantity:  defi.TokenAmountFloatToWEI(10.0, defi.TokenSTG),
			FromToken: defi.TokenSTG,
			ToToken:   defi.TokenSTG,
		})
		assert.NoError(t, err)

		println(len(tx.Tx))

	})

	t.Run("transfer", func(t *testing.T) {
		ctx := context.Background()

		wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
		assert.NoError(t, err)

		b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
			WalletAddress: tests.GetConfig().Wallet,
			Token:         defi.TokenUSDC,
		})

		println(b.HumanReadable)

		v := defi.TokenAmountFloatToWEI(0.5, defi.TokenUSDC)

		res, err := r.Transfer(ctx, &defi.TransferReq{
			Wallet: wallet,
			ToAddr: common.HexToAddress(""),
			Token:  defi.TokenUSDC,
			Amount: v,
		})
		assert.NoError(t, err)
		println(res.Tx.Hash().String())
	})

}
