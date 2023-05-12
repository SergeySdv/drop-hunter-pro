package arbitrum

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
	MainNetURL = "https://arb1.arbitrum.io/rpc"
)

// docs https://arbiscan.io/tokens?p=1
var TokenAddress = map[defi.Token]common.Address{
	v1.Token_USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
	v1.Token_STG:  common.HexToAddress("0x6694340fc020c5e6b96567843da2df01b2ce1eb6"),
	v1.Token_USDC: common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8"),
}

var Dict = defi.Dict{
	Stargate: defi.Stargate{
		STGStakingAddress:        common.HexToAddress("0xfbd849e6007f9bc3cc2d6eb159c045b8dc660268"),
		StargateRouterAddress:    common.HexToAddress("0x53Bf833A5d6c4ddA888F69c22C88C9f356a41614"),
		StargateRouterEthAddress: common.HexToAddress("0xbf22f0f184bCcbeA268dF387a49fF5238dD23E40"),
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
	return "https://arbiscan.io/tx/" + txId
}

func NewClient(c *ClientConfig) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	ethcli, err := defi.NewEtheriumClient(&defi.ClientConfig{
		Network:   v1.Network_ARBITRUM,
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
