package v1

import (
	"context"
	"crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	settings2 "crypto_scripts/internal/server/settings"
	"crypto_scripts/internal/server/user"
	"crypto_scripts/internal/uniclient"
)

type SettingsService struct {
	v1.UnimplementedSettingsServiceServer
	rep repository.SettingsRepository
}

func NewSettingsService(rep repository.SettingsRepository) *SettingsService {
	return &SettingsService{rep: rep}
}

func (s *SettingsService) Reset(ctx context.Context, req *v1.ResetRequest) (*v1.ResetResponse, error) {

	defaultSettings, err := settings2.DefaultSettings(ctx, user.GetUserId(ctx))
	if err != nil {
		return nil, err
	}

	if err := s.rep.UpdateSettings(ctx, defaultSettings); err != nil {
		return nil, err
	}
	settings, err := s.rep.GetSettings(ctx, user.GetUserId(ctx))
	if err != nil {
		return nil, err
	}

	return &v1.ResetResponse{
		Settings: settings,
	}, nil
}

func (s *SettingsService) GetNetworkByRPC(ctx context.Context, req *v1.GetNetworkByRPCRequest) (*v1.GetNetworkByRPCResponse, error) {
	client, err := uniclient.NewBaseClient(req.Network, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: req.Endpoint,
	})
	if err != nil {
		e := "invalid rpc"
		return &v1.GetNetworkByRPCResponse{
			Valid:             false,
			SuggestedGasPrice: 0,
			Id:                0,
			Error:             &e,
		}, nil
	}

	networkId := client.GetNetworkId()
	if err != nil {
		e := "error getting network id"
		return &v1.GetNetworkByRPCResponse{
			Valid:             false,
			SuggestedGasPrice: 0,
			Id:                0,
			Error:             &e,
		}, nil
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		e := "error getting gas price"
		return &v1.GetNetworkByRPCResponse{
			Valid:             false,
			SuggestedGasPrice: 0,
			Id:                0,
			Error:             &e,
		}, nil
	}

	return &v1.GetNetworkByRPCResponse{
		Valid:             true,
		SuggestedGasPrice: gasPrice.Int64(),
		Id:                networkId.Int64(),
	}, nil
}
func (s *SettingsService) GetSettings(ctx context.Context, req *v1.GetSettingsRequest) (*v1.GetSettingsResponse, error) {

	settings, err := settings2.ResolveSettingsForUser(ctx, user.GetUserId(ctx), s.rep)
	if err != nil {
		return nil, err
	}

	return &v1.GetSettingsResponse{
		Settings: settings,
	}, nil

}
func (s *SettingsService) UpdateSettings(ctx context.Context, req *v1.UpdateSettingsRequest) (*v1.UpdateSettingsResponse, error) {

	if err := s.rep.UpdateSettings(ctx, req.Settings); err != nil {
		return nil, err
	}

	settings, err := s.rep.GetSettings(ctx, user.GetUserId(ctx))
	if err != nil {
		return nil, err
	}

	return &v1.UpdateSettingsResponse{
		Settings: settings,
	}, nil
}
