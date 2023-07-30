package zksyncera

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestVelocore(t *testing.T) {
	r, err := NewMainNetClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	ctx := context.Background()

	//data, err := r.makeMaverickSwapData(ctx, v1.Token_USDC, v1.Token_ETH, tests.GetConfig().Wallet, big.NewInt(10*10e5))
	//assert.NoError(t, err)
	//println("0x" + common.Bytes2Hex(data))

	t.Run("ETH -> USDC", func(t *testing.T) {
		//t.Skip()
		am := defi.TokenAmountFloatToWEI(0.001, v1.Token_ETH)

		req := &defi.DefaultSwapReq{
			Network:   v1.Network_ZKSYNCERA,
			Amount:    am,
			FromToken: v1.Token_ETH,
			ToToken:   v1.Token_USDC,
			WalletPK:  tests.GetConfig().PK,
			Debug:     true,
		}

		data, err := r.makeSpaceFiSwapData(ctx, req)
		assert.NoError(t, err)
		println("have: " + "0x" + common.Bytes2Hex(data))
		expected := "0x7ff36ab50000000000000000000000000000000000000000000000000000000000d0a69300000000000000000000000000000000000000000000000000000000000000800000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150000000000000000000000000000000000000000000000000000000064c314f700000000000000000000000000000000000000000000000000000000000000020000000000000000000000005aea5775959fbc2557cc8789bc1bf90a239d9a910000000000000000000000003355df6d4c9c3035724fd0e3914de96a5a83aaf4"
		println("want: " + expected)
	})

	t.Run("USDC -> ETH", func(t *testing.T) {
		//t.Skip()
		am := defi.TokenAmountFloatToWEI(10, v1.Token_USDC)

		req := &defi.DefaultSwapReq{
			Network:   v1.Network_ZKSYNCERA,
			Amount:    am,
			FromToken: v1.Token_USDC,
			ToToken:   v1.Token_ETH,
			WalletPK:  tests.GetConfig().PK,
			Debug:     true,
		}

		data, err := r.makeSpaceFiSwapData(ctx, req)
		assert.NoError(t, err)
		println("have: " + "0x" + common.Bytes2Hex(data))
		expected := "0x18cbafe500000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000001e3c72cf0d4f000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150000000000000000000000000000000000000000000000000000000064c31cf700000000000000000000000000000000000000000000000000000000000000020000000000000000000000003355df6d4c9c3035724fd0e3914de96a5a83aaf40000000000000000000000005aea5775959fbc2557cc8789bc1bf90a239d9a91"
		println("want: " + expected)
	})

}
