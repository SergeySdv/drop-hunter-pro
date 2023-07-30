package v1

import (
	"context"
	"strconv"

	"github.com/hardstylez72/cry/internal/go1inch"
	"github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type Swap1inchService struct {
	v1.UnimplementedSwap1InchServiceServer
	swap1inch *go1inch.Client
}

func NewSwap1inchService() *Swap1inchService {
	return &Swap1inchService{
		swap1inch: go1inch.NewClient(&go1inch.Config{}),
	}
}

var Swap1inchNetworkMap = map[go1inch.Network]v1.Network{
	go1inch.Arbitrum:    v1.Network_ARBITRUM,
	go1inch.Optimism:    v1.Network_OPTIMISM,
	go1inch.Avalanche:   v1.Network_AVALANCHE,
	go1inch.Auror:       -1,
	go1inch.Bsc:         v1.Network_BinanaceBNB,
	go1inch.Eth:         v1.Network_Etherium,
	go1inch.Fantom:      -1,
	go1inch.Klaytn:      -1,
	go1inch.GnosisChain: -1,
	go1inch.Matic:       v1.Network_POLIGON,
	go1inch.ZkSync:      v1.Network_ZKSYNCERA,
}

var Swap1inchNetworkMapReverse = map[v1.Network]go1inch.Network{
	v1.Network_ARBITRUM:    go1inch.Arbitrum,
	v1.Network_OPTIMISM:    go1inch.Optimism,
	v1.Network_AVALANCHE:   go1inch.Avalanche,
	v1.Network_BinanaceBNB: go1inch.Bsc,
	v1.Network_Etherium:    go1inch.Eth,
	v1.Network_POLIGON:     go1inch.Matic,
	v1.Network_ZKSYNCERA:   go1inch.ZkSync,
}

func (s *Swap1inchService) GetNetworks(ctx context.Context, req *v1.GetNetworksRequest) (*v1.GetNetworksResponse, error) {

	all := go1inch.GetNetworks()

	result := make([]v1.Network, 0)

	for _, network := range all {
		n, supported := Swap1inchNetworkMap[network]
		if !supported {
			continue
		}
		if n == -1 {
			continue
		}

		result = append(result, n)
	}

	return &v1.GetNetworksResponse{
		Networks: result,
	}, nil
}

func (s *Swap1inchService) GetTokens(ctx context.Context, req *v1.GetTokensRequest) (*v1.GetTokensResponse, error) {

	network, supported := Swap1inchNetworkMapReverse[req.Network]
	if !supported {
		return nil, errors.New("network is not supported")
	}

	res, _, err := s.swap1inch.Tokens(ctx, network)
	if err != nil {
		return nil, err
	}

	result := make([]*v1.Swap1InchToken, 0)

	for _, token := range res.Tokens {
		result = append(result, &v1.Swap1InchToken{
			Code: token.Symbol,
			Name: token.Name,
			Addr: token.Address,
		})
	}

	return &v1.GetTokensResponse{Tokens: result}, nil
}

func (s *Swap1inchService) GetSwapOptions(ctx context.Context, req *v1.GetSwapOptionsRequest) (*v1.GetSwapOptionsResponse, error) {
	network, supported := Swap1inchNetworkMapReverse[req.Network]
	if !supported {
		return nil, errors.New("network is not supported")
	}

	res, _, err := s.swap1inch.Quote(ctx, network, req.FromTokenAddr, req.ToTokenAddr, req.Amount, nil)
	if err != nil {
		return nil, err
	}

	result := make([]*v1.GetSwapOptionsProtocol, 0)

	for _, protocols := range res.Protocols {
		for _, pcs := range protocols {
			for _, p := range pcs {
				result = append(result, &v1.GetSwapOptionsProtocol{
					Name:         p.Name,
					EstimatedGas: strconv.Itoa(int(res.EstimatedGas)),
				})
			}

		}
	}

	return &v1.GetSwapOptionsResponse{
		Protocols: result,
	}, nil
}
