package defi

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

const RetryMax = 1

type Dict struct {
	Stargate                 Stargate
	TestNetBridgeSwapAddress common.Address
}

type SyncSwap struct {
	RouterSwap         common.Address
	ClassicPoolFactory common.Address
}

type Stargate struct {
	StargateRouterAddress    common.Address
	StargateRouterEthAddress common.Address
}

type EtheriumClient struct {
	Cli    *ethclient.Client
	Cfg    *ClientConfig
	RpcCli *rpc.Client
}

type ClientConfig struct {
	Network   Network
	MainToken Token
	MainNet   string
	TokenMap  map[Token]common.Address
	Dict      *Dict
	Httpcli   *http.Client
	TxViewFn  func(s string) string
	networkId *big.Int
}

func NewEVMClient(c *ClientConfig) (*EtheriumClient, error) {

	rpcClient, err := rpc.DialOptions(context.Background(), c.MainNet, rpc.WithHTTPClient(c.Httpcli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.MainNet)
	}

	return &EtheriumClient{
		Cli:    ethclient.NewClient(rpcClient),
		Cfg:    c,
		RpcCli: rpcClient,
	}, nil
}
