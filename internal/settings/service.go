package settings

import (
	"context"
	"sync"

	"github.com/hardstylez72/cry/internal/defi/arbitrum"
	"github.com/hardstylez72/cry/internal/defi/avalanche"
	"github.com/hardstylez72/cry/internal/defi/bnb"
	"github.com/hardstylez72/cry/internal/defi/etherium"
	"github.com/hardstylez72/cry/internal/defi/optimism"
	"github.com/hardstylez72/cry/internal/defi/poligon"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	"github.com/hardstylez72/cry/internal/defi/zksynclite"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type GetSettingsNetworkRequest struct {
	Network v1.Network
	UserId  string
}

type Service struct {
	rep repository.SettingsRepository
}

func NewService(rep repository.SettingsRepository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetSettings(ctx context.Context, userId string) (*v1.Settings, error) {
	return s.rep.GetSettings(ctx, userId)
}

func (s *Service) UpdateSettings(ctx context.Context, in *v1.Settings) error {
	return s.rep.UpdateSettings(ctx, in)
}

func (s *Service) GetSettingsNetwork(ctx context.Context, req *GetSettingsNetworkRequest) (*v1.SettingsNetwork, error) {

	stgs, err := s.ResolveSettingsForUser(ctx, req.UserId)
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
	case v1.Network_ZKSYNCERA:
		return stgs.ZksyncMainNet, nil
	case v1.Network_ZKSYNCERATESTNET:
		return stgs.ZksyncTestNet, nil
	case v1.Network_ZKSYNCLITE:
		return stgs.ZksyncLite, nil
	}

	return nil, errors.New("usupported network: " + req.Network.String())
}

func GetSettingsNetworkSource(network v1.Network, stgs *v1.Settings) (*v1.SettingsNetwork, error) {
	switch network {
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
	case v1.Network_ZKSYNCERATESTNET:
		return stgs.ZksyncTestNet, nil
	case v1.Network_ZKSYNCERA:
		return stgs.ZksyncMainNet, nil
	case v1.Network_ZKSYNCLITE:
		return stgs.ZksyncLite, nil
	default:
		return nil, errors.New("unknown network: " + network.String())
	}
}

var mu sync.Mutex = sync.Mutex{}

func (s *Service) ResolveSettingsForUser(ctx context.Context, userId string) (*v1.Settings, error) {
	mu.Lock()
	defer mu.Unlock()

	rep := s.rep

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

	// new networks
	if settings.ZksyncTestNet == nil || settings.ZksyncMainNet == nil {
		zksyncTestNet, err := uniclient.NewBaseClient(v1.Network_ZKSYNCERATESTNET, &uniclient.BaseClientConfig{
			RPCEndpoint: zksyncera.TestNetURL,
		})
		if err != nil {
			return nil, err
		}

		zksyncCli, err := uniclient.NewBaseClient(v1.Network_ZKSYNCERA, &uniclient.BaseClientConfig{
			RPCEndpoint: zksyncera.MainNetURL,
		})
		if err != nil {
			return nil, err
		}

		settings.ZksyncTestNet = &v1.SettingsNetwork{
			Id:          zksyncTestNet.GetNetworkId().Int64(),
			RpcEndpoint: zksyncera.TestNetURL,
		}

		settings.ZksyncMainNet = &v1.SettingsNetwork{
			Id:          zksyncCli.GetNetworkId().Int64(),
			RpcEndpoint: zksyncera.MainNetURL,
		}

		if err := rep.UpdateSettings(ctx, settings); err != nil {
			return rep.GetSettings(ctx, userId)
		}
		return rep.GetSettings(ctx, userId)
	}

	if settings.ZksyncLite == nil {

		zksyncCli, err := uniclient.NewBaseClient(v1.Network_ZKSYNCLITE, &uniclient.BaseClientConfig{
			RPCEndpoint: zksynclite.MainNetURL,
		})
		if err != nil {
			return nil, err
		}

		settings.ZksyncLite = &v1.SettingsNetwork{
			Id:          zksyncCli.GetNetworkId().Int64(),
			RpcEndpoint: zksynclite.MainNetURL,
		}

		if err := rep.UpdateSettings(ctx, settings); err != nil {
			return rep.GetSettings(ctx, userId)
		}
		return rep.GetSettings(ctx, userId)
	}

	eth := "10000000000000000"
	poligon := "40000000000000000000"
	avalanche := "1000000000000000000"
	bnb := "100000000000000000"
	if settings.ZksyncTestNet.MaxGas == nil {
		settings.ZksyncTestNet.MaxGas = &eth
	}
	if settings.ZksyncMainNet.MaxGas == nil {
		settings.ZksyncMainNet.MaxGas = &eth
	}
	if settings.ZksyncLite.MaxGas == nil {
		settings.ZksyncLite.MaxGas = &eth
	}
	if settings.Arbitrum.MaxGas == nil {
		settings.Arbitrum.MaxGas = &eth
	}
	if settings.Optimism.MaxGas == nil {
		settings.Optimism.MaxGas = &eth
	}
	if settings.Etherium.MaxGas == nil {
		settings.Etherium.MaxGas = &eth
	}
	if settings.Polygon.MaxGas == nil {
		settings.Polygon.MaxGas = &poligon
	}
	if settings.Avalanche.MaxGas == nil {
		settings.Avalanche.MaxGas = &avalanche
	}
	if settings.Bnb.MaxGas == nil {
		settings.Bnb.MaxGas = &bnb
	}

	return settings, nil
}

func DefaultSettings(ctx context.Context, userId string) (*v1.Settings, error) {

	etheriumCli, err := uniclient.NewBaseClient(v1.Network_Etherium, &uniclient.BaseClientConfig{
		RPCEndpoint: etherium.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_Etherium ok")

	bnbCli, err := uniclient.NewBaseClient(v1.Network_BinanaceBNB, &uniclient.BaseClientConfig{
		RPCEndpoint: bnb.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_BinanaceBNB ok")

	poligonCli, err := uniclient.NewBaseClient(v1.Network_POLIGON, &uniclient.BaseClientConfig{
		RPCEndpoint: poligon.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_POLIGON ok")

	arbitrumCli, err := uniclient.NewBaseClient(v1.Network_ARBITRUM, &uniclient.BaseClientConfig{
		RPCEndpoint: arbitrum.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_ARBITRUM ok")

	optimismCli, err := uniclient.NewBaseClient(v1.Network_OPTIMISM, &uniclient.BaseClientConfig{
		RPCEndpoint: optimism.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_OPTIMISM ok")

	avalancheCli, err := uniclient.NewBaseClient(v1.Network_AVALANCHE, &uniclient.BaseClientConfig{
		RPCEndpoint: avalanche.MainNetURL,
	})
	if err != nil {
		return nil, err
	}
	log.Log.Debug("settings v1.Network_AVALANCHE ok")

	zksyncTestNet, err := uniclient.NewBaseClient(v1.Network_ZKSYNCERATESTNET, &uniclient.BaseClientConfig{
		RPCEndpoint: zksyncera.TestNetURL,
	})
	if err != nil {
		return nil, err
	}

	zksyncCli, err := uniclient.NewBaseClient(v1.Network_ZKSYNCERA, &uniclient.BaseClientConfig{
		RPCEndpoint: zksyncera.MainNetURL,
	})
	if err != nil {
		return nil, err
	}

	zksyncLiteCli, err := uniclient.NewBaseClient(v1.Network_ZKSYNCLITE, &uniclient.BaseClientConfig{
		RPCEndpoint: zksynclite.MainNetURL,
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
			Id:          poligonCli.GetNetworkId().Int64(),
			RpcEndpoint: poligon.MainNetURL,
		},
		Avalanche: &v1.SettingsNetwork{
			Id:          avalancheCli.GetNetworkId().Int64(),
			RpcEndpoint: avalanche.MainNetURL,
		},
		ZksyncTestNet: &v1.SettingsNetwork{
			Id:          zksyncTestNet.GetNetworkId().Int64(),
			RpcEndpoint: zksyncera.TestNetURL,
		},
		ZksyncMainNet: &v1.SettingsNetwork{
			Id:          zksyncCli.GetNetworkId().Int64(),
			RpcEndpoint: zksyncera.MainNetURL,
		},
		ZksyncLite: &v1.SettingsNetwork{
			Id:          zksyncLiteCli.GetNetworkId().Int64(),
			RpcEndpoint: zksynclite.MainNetURL,
		},
	}

	out.TaskGasLimitMap = map[string]int64{
		v1.TaskType_StargateBridge.String():                 2000000000000000,
		v1.TaskType_ZkSyncOfficialBridgeToEthereum.String(): 2000000000000000,
		v1.TaskType_SyncSwap.String():                       2000000000000000,
	}
	return out, nil
}
