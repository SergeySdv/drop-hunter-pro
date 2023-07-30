package zksynclite

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	c, err := NewClient(&Config{
		HttpCli:   tests.GetConfig().Cli,
		RPCETHURL: "https://api.zksync.io/jsrpc",
	})
	assert.NotNil(t, c)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//b, err := c.GetBalance(ctx, &defi.GetBalanceReq{
	//	WalletAddress: tests.GetConfig().Wallet,
	//	Token:         v1.Token_ETH,
	//})
	//
	//assert.NoError(t, err)
	//assert.NotNil(t, b)

	tx, err := c.TxInfo(ctx, common.HexToHash("0x527b4d179355f5555e75abd1b5786bf14a2b57bcc387bd8318216efc743c224a"))
	assert.NoError(t, err)
	assert.NotNil(t, tx)
}
