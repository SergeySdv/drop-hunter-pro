package arbitrum

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	c := tests.GetConfig()

	wallet, err := defi.NewWalletTransactor(c.PK)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
		WalletAddress: wallet.WalletAddr,
		Token:         v1.Token_ETH,
	})
	assert.NoError(t, err)
	assert.NotNil(t, b)
	//res, err := r.Transfer(context.Background(), &defi.TransferReq{
	//	Wallet: wallet,
	//	ToAddr: common.HexToAddress("0xb83f35d9d80ceff4d9ad93eff941d5aba64a6341"),
	//	Token:  v1.Token_ETH,
	//	Amount: b.WEI,
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, res)
	//println(r.TxViewFn(res.Tx.Hash().String()))

	//logs, err := r.defi.Cli.FilterLogs(context.Background(), ethereum.FilterQuery{
	//	BlockHash: nil,
	//	FromBlock: nil,
	//	ToBlock:   nil,
	//	Addresses: []common.Address{common.HexToAddress("0x7c7273AA418074F8a9209eC901474AE6C60FA231")},
	//	Topics:    nil,
	//})
	//
	//assert.NoError(t, err)
	//assert.NotNil(t, logs)
	//tx, _, _ := r.defi.Cli.TransactionByHash(context.Background(), common.HexToHash("0x2693486b5ce16b8d6514c899158470df323f061ba131506559c1ddce39710e19"))

	// gas limit 1655806
	// gas fee cap 200000000
	// gas tip cap 0
	// value 248791766200233
	//assert.NotNil(t, tx)
	//println(1)
	//h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	//assert.NoError(t, err)
	//assert.NotNil(t, h)

	//t.Run("balance", func(t *testing.T) {
	//	t.Skip()
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         defi.TokenETH,
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
	//	am := defi.TokenAmountToWEI(6, defi.TokenSTG)
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
	//	t.Skip()
	//	wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
	//	assert.NoError(t, err)
	//
	//	tx, err := r.StargateBridgeSwap(ctx, &defi.StargateBridgeSwapReq{
	//		DestChain: defi.SBNameOptimism,
	//		Wallet:    wallet,
	//		Quantity:  defi.TokenAmountFloatToWEI(10.0, defi.TokenSTG),
	//		FromToken: defi.TokenSTG,
	//		ToToken:   defi.TokenSTG,
	//	})
	//	assert.NoError(t, err)
	//
	//	println(len(tx.Tx))
	//
	//})
	//
	//t.Run("transfer", func(t *testing.T) {
	//	ctx := context.Background()
	//
	//	wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
	//	assert.NoError(t, err)
	//
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         defi.TokenUSDC,
	//	})
	//
	//	println(b.HumanReadable)
	//
	//	v := defi.TokenAmountFloatToWEI(0.5, defi.TokenUSDC)
	//
	//	res, err := r.Transfer(ctx, &defi.TransferReq{
	//		Wallet: wallet,
	//		ToAddr: common.HexToAddress(""),
	//		Token:  defi.TokenUSDC,
	//		Amount: v,
	//	})
	//	assert.NoError(t, err)
	//	println(res.Tx.Hash().String())
	//})

}
