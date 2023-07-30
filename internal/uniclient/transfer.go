package uniclient

import (
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/arbitrum"
	"github.com/hardstylez72/cry/internal/defi/avalanche"
	"github.com/hardstylez72/cry/internal/defi/bnb"
	"github.com/hardstylez72/cry/internal/defi/etherium"
	"github.com/hardstylez72/cry/internal/defi/optimism"
	"github.com/hardstylez72/cry/internal/defi/poligon"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewTransfer(network v1.Network, c *BaseClientConfig) (defi.Transfer, error) {

	proxy, err := socks5.NewSock5ProxyString(c.ProxyString, c.UserAgentHeader)
	if err != nil {
		return nil, err
	}

	var cli defi.Transfer
	switch network {
	case v1.Network_ARBITRUM:
		cli, err = arbitrum.NewClient(&arbitrum.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_Etherium:
		cli, err = etherium.NewClient(&etherium.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_BinanaceBNB:
		cli, err = bnb.NewClient(&bnb.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_OPTIMISM:
		cli, err = optimism.NewClient(&optimism.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_POLIGON:
		cli, err = poligon.NewClient(&poligon.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	case v1.Network_AVALANCHE:
		cli, err = avalanche.NewClient(&avalanche.ClientConfig{HttpCli: proxy.Cli, RPCEndpoint: c.RPCEndpoint})
	default:
		return nil, errors.New("network is not supported for Transfer")
	}

	if err != nil {
		return nil, err
	}

	return cli, nil
}
