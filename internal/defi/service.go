package defi

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

const RetryMax = 1

type Dict struct {
	Stargate Stargate
}

type Stargate struct {
	STGStakingAddress        common.Address
	StargateRouterAddress    common.Address
	StargateRouterEthAddress common.Address
}

type EtheriumClient struct {
	cli *ethclient.Client
	c   *ClientConfig
}

type ClientConfig struct {
	Network   Network
	MainToken Token
	MainNet   string
	TokenMap  map[Token]common.Address
	Dict      *Dict
	Httpcli   *http.Client
	TxViewFn  func(s string) string
	GasLimit  *GasLimit
	networkId *big.Int
}

func NewEtheriumClient(c *ClientConfig) (*EtheriumClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rpcClient, err := rpc.DialOptions(ctx, c.MainNet, rpc.WithHTTPClient(c.Httpcli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.MainNet)
	}
	ethcli := ethclient.NewClient(rpcClient)

	return &EtheriumClient{
		cli: ethcli,
		c:   c,
	}, nil
}
