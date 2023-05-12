package bnb

import (
	"context"
	"crypto_scripts/internal/defi"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
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
			Token:         v1.Token_USDT,
		})
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})
}
