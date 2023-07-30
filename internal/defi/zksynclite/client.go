package zksynclite

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type Client struct {
	ethcli *ethclient.Client
	Cfg    *Config
}

type Config struct {
	HttpCli   *http.Client
	RPCETHURL string

	network   v1.Network
	mainToken v1.Token
	tokenMap  map[v1.Token]common.Address
	txViewFn  func(s string) string
	networkId *big.Int
}

const (
	MainNetURL = "https://api.zksync.io/jsrpc" // mainnet
)

func NewClient(c *Config) (*Client, error) {

	rpcClient, err := rpc.DialOptions(context.Background(), c.RPCETHURL, rpc.WithHTTPClient(c.HttpCli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.RPCETHURL)
	}

	c.network = v1.Network_ZKSYNCLITE
	c.mainToken = v1.Token_ETH
	c.networkId = big.NewInt(300)

	return &Client{
		ethcli: ethclient.NewClient(rpcClient),
		Cfg:    c,
	}, nil
}

func (c *Client) TxViewFn(id string) string {
	return "https://zkscan.io/explorer/transactions/" + id
}
func (c *Client) GetNetworkId() *big.Int {
	return c.Cfg.networkId
}
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return nil, errors.New("not implemented") // todo:...
}
