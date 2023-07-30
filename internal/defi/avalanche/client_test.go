package avalanche

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: "https://rpc.ankr.com/avalanche",
		GasLimit:    nil,
	})

	h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, h)

	//assert.NoError(t, err)
	//assert.NotNil(t, r)
	//
	//t.Run("balance", func(t *testing.T) {
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         v1.Token_MATIC,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})

}
