package bnb

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

	//MainNetURL https://docs.polygonscan.com/getting-started/endpoint-urls
	//MainNetURL = "https://bsc-dataseed.binance.org"
	MainNetURL = "https://rpc.ankr.com/bsc"
)

var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT: common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	v1.Token_STG:  common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b"),
	v1.Token_USDC: common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		STGStakingAddress:        common.HexToAddress("not supported"),
		StargateRouterAddress:    common.HexToAddress("0x4a364f8c717cAAD9A442737Eb7b8A55cc6cf18D8"),
		StargateRouterEthAddress: common.HexToAddress("not supported"),
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

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEtheriumClient(&defi.ClientConfig{
		Network:   v1.Network_BinanaceBNB,
		MainToken: v1.Token_BNB,
		MainNet:   c.RPCEndpoint,
		TokenMap:  TokenAddress,
		Dict:      &Dict,
		Httpcli:   config.HttpCli,
		TxViewFn: func(txId string) string {
			return "https://bscscan.com/tx/" + txId
		},
		GasLimit: c.GasLimit,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ethereum net: "+c.RPCEndpoint)
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
