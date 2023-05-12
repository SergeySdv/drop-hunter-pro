package uniclient

import (
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/defi/arbitrum"
	"crypto_scripts/internal/defi/avalanche"
	"crypto_scripts/internal/defi/bnb"
	"crypto_scripts/internal/defi/etherium"
	"crypto_scripts/internal/defi/optimism"
	"crypto_scripts/internal/defi/poligon"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/socks5"

	"github.com/pkg/errors"
)

type BaseClientConfig struct {
	ProxyString string
	RPCEndpoint string
	GasLimit    *defi.GasLimit
}

func NewBaseClient(network v1.Network, c *BaseClientConfig) (defi.Networker, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString)
	if err != nil {
		return nil, err
	}

	var cli defi.Networker
	switch network {
	case v1.Network_ARBITRUM:
		cli, err = arbitrum.NewClient(&arbitrum.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	case v1.Network_Etherium:
		cli, err = etherium.NewClient(&etherium.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	case v1.Network_BinanaceBNB:
		cli, err = bnb.NewClient(&bnb.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	case v1.Network_OPTIMISM:
		cli, err = optimism.NewClient(&optimism.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	case v1.Network_POLIGON:
		cli, err = poligon.NewClient(&poligon.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	case v1.Network_AVALANCHE:
		cli, err = avalanche.NewClient(&avalanche.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint, GasLimit: c.GasLimit})
	default:
		return nil, errors.New("network is not supported for Networker")
	}
	return cli, err
}
