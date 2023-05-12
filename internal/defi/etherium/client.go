package etherium

import (
	"context"
	"crypto_scripts/internal/defi"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

const (
	MainNetURL = "https://cloudflare-eth.com"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
	v1.Token_STG:  common.HexToAddress("0xaf5191b0de278c7286d6c7cc6ab6bb8a73ba2cd6"),
	v1.Token_USDC: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		STGStakingAddress:        common.HexToAddress("0x0e42acbd23faee03249daff896b78d7e79fbd58e"),
		StargateRouterAddress:    common.HexToAddress("0x8731d54E9D02c286767d56ac03e8037C07e01e98"),
		StargateRouterEthAddress: common.HexToAddress("0x150f94B44927F078737562f0fcF3C95c01Cc2376"),
	},
}

type Client struct {
	defi      *defi.EtheriumClient
	NetworkId *big.Int
}

type ClientConfig struct {
	HttpCli     *http.Client
	RPCEndpoint string
	GasLimit    *defi.GasLimit
}

func TxViewer(txId string) string {
	return "https://etherscan.io/token/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {
	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}
	ethcli, err := defi.NewEtheriumClient(&defi.ClientConfig{
		Network:   v1.Network_Etherium,
		MainToken: v1.Token_ETH,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      &Dict,
		Httpcli:   config.HttpCli,
		TxViewFn:  TxViewer,
		GasLimit:  c.GasLimit,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum main: "+c.RPCEndpoint)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	networkId, err := ethcli.GetNetworkId(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		defi:      ethcli,
		NetworkId: networkId,
	}, nil
}
