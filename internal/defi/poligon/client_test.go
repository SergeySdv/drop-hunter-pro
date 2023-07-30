package poligon

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	r.WaitTxComplete(context.Background(), common.HexToHash("0xf521afb3009e2d90ac54527856b788496fd260a06e05cb24d96e9b1548d33208"))

	h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, h)
	t.Run("balance", func(t *testing.T) {
		b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
			WalletAddress: tests.GetConfig().Wallet,
			Token:         v1.Token_MATIC,
		})
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

}
