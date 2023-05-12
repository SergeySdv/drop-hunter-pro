package etherium

import (
	"context"
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	t.Run("balance", func(t *testing.T) {
		b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
			WalletAddress: tests.GetConfig().Wallet,
			Token:         defi.TokenETH,
		})
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	t.Run("stg stake", func(t *testing.T) {

		wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
		assert.NoError(t, err)

		am := defi.TokenAmountToWEI(6)

		res, err := r.StakeSTG(context.Background(), &defi.StakeSTGReq{
			Amount: am,
			Wallet: wallet,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
