package bnb

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, h)

	//tx, penging, err := r.defi.Cli.TransactionByHash(context.Background(), common.HexToHash("0x02433465f2e9a56484dea517b4ab932ca1d5970451ea8cef90499713eb447021"))
	//assert.NoError(t, err)
	//assert.NotNil(t, penging)
	//assert.NotNil(t, tx)
	//rec, err := bind.WaitMined(context.Background(), r.defi.Cli, tx)
	//assert.NotNil(t, rec)
	//assert.NotNil(t, err)
	//t.Run("balance", func(t *testing.T) {
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         v1.Token_USDT,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})
}
