package settings

import (
	"context"
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/defi/arbitrum"
	"crypto_scripts/internal/defi/avalanche"
	"crypto_scripts/internal/defi/bnb"
	"crypto_scripts/internal/defi/etherium"
	"crypto_scripts/internal/defi/optimism"
	"crypto_scripts/internal/defi/poligon"
	"crypto_scripts/internal/log"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/uniclient"
	"math/big"
	"sync"

	"github.com/pkg/errors"
)

type GetSettingsNetworkRequest struct {
	Network v1.Network
	UserId  string
	Rep     repository.SettingsRepository
}

func GetSettingsNetwork(ctx context.Context, req *GetSettingsNetworkRequest) (*v1.SettingsNetwork, error) {

	stgs, err := ResolveSettingsForUser(ctx, req.UserId, req.Rep)
	if err != nil {
		return nil, err
	}
	switch req.Network {
	case v1.Network_BinanaceBNB:
		return stgs.Bnb, nil
	case v1.Network_ARBITRUM:
		return stgs.Arbitrum, nil
	case v1.Network_OPTIMISM:
		return stgs.Optimism, nil
	case v1.Network_Etherium:
		return stgs.Etherium, nil
	case v1.Network_POLIGON:
		return stgs.Polygon, nil
	case v1.Network_AVALANCHE:
		return stgs.Avalanche, nil
	}

	return nil, errors.New("usupported network: " + req.Network.String())
}

func GetGasLimit(network *v1.SettingsNetwork) *defi.GasLimit {
	if network == nil {
		return nil
	}

	if !network.UseLimitGas {
		return nil
	}

	GasTotal := big.NewInt(0)
	if network.GasTotal != nil {
		GasTotal = big.NewInt(network.GetGasTotal())
	}

	return &defi.GasLimit{
		TotalGas: GasTotal,
	}
}

var mu sync.Mutex = sync.Mutex{}

func ResolveSettingsForUser(ctx context.Context, userId string, rep repository.SettingsRepository) (*v1.Settings, error) {
	mu.Lock()
	defer mu.Unlock()
	settings, err := rep.GetSettings(ctx, userId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {

			s, err := DefaultSettings(ctx, userId)
			if err != nil {
				return nil, err
			}

			if err := rep.UpdateSettings(ctx, s); err != nil {
				return rep.GetSettings(ctx, userId)
			}
		}
		return nil, err
	}
	return settings, nil
}

func DefaultSettings(ctx context.Context, userId string) (*v1.Settings, error) {

	etheriumCli, err := uniclient.NewBaseClient(v1.Network_Etherium, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: etherium.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_Etherium ok")

	bnbCli, err := uniclient.NewBaseClient(v1.Network_BinanaceBNB, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: bnb.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_BinanaceBNB ok")

	//poligonCli, err := uniclient.NewBaseClient(v1.Network_POLIGON, &uniclient.BaseClientConfig{
	//	ProxyString: "",
	//	RPCEndpoint: poligon.MainNetURL,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//log.Log.Debug("settings v1.Network_POLIGON ok")

	arbitrumCli, err := uniclient.NewBaseClient(v1.Network_ARBITRUM, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: arbitrum.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_ARBITRUM ok")

	optimismCli, err := uniclient.NewBaseClient(v1.Network_OPTIMISM, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: optimism.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_OPTIMISM ok")

	avalancheCli, err := uniclient.NewBaseClient(v1.Network_AVALANCHE, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: avalanche.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_AVALANCHE ok")

	out := &v1.Settings{
		UserId: userId,
		Bnb: &v1.SettingsNetwork{
			Id:          bnbCli.GetNetworkId().Int64(),
			RpcEndpoint: bnb.MainNetURL,
		},
		Optimism: &v1.SettingsNetwork{
			Id:          optimismCli.GetNetworkId().Int64(),
			RpcEndpoint: optimism.MainNetURL,
		},
		Arbitrum: &v1.SettingsNetwork{
			Id:          arbitrumCli.GetNetworkId().Int64(),
			RpcEndpoint: arbitrum.MainNetURL,
		},
		Etherium: &v1.SettingsNetwork{
			Id:          etheriumCli.GetNetworkId().Int64(),
			RpcEndpoint: etherium.MainNetURL,
		},
		Polygon: &v1.SettingsNetwork{
			//Id:          poligonCli.GetNetworkId().Int64(),
			RpcEndpoint: poligon.MainNetURL,
		},
		Avalanche: &v1.SettingsNetwork{
			Id:          avalancheCli.GetNetworkId().Int64(),
			RpcEndpoint: avalanche.MainNetURL,
		},
	}

	return out, nil
}
