package v1

import (
	"context"

	"github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/user"
	settings2 "github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/uniclient"
)

type SettingsService struct {
	v1.UnimplementedSettingsServiceServer
	settingsService *settings2.Service
}

func NewSettingsService(settingsService *settings2.Service) *SettingsService {
	return &SettingsService{settingsService: settingsService}
}

func (s *SettingsService) Reset(ctx context.Context, req *v1.ResetRequest) (*v1.ResetResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	defaultSettings, err := settings2.DefaultSettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err := s.settingsService.UpdateSettings(ctx, defaultSettings); err != nil {
		return nil, err
	}
	settings, err := s.settingsService.GetSettings(ctx, userId)
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

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	settings, err := s.settingsService.ResolveSettingsForUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetSettingsResponse{
		Settings: settings,
	}, nil

}
func (s *SettingsService) UpdateSettings(ctx context.Context, req *v1.UpdateSettingsRequest) (*v1.UpdateSettingsResponse, error) {

	if err := s.settingsService.UpdateSettings(ctx, req.Settings); err != nil {
		return nil, err
	}

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	settings, err := s.settingsService.GetSettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateSettingsResponse{
		Settings: settings,
	}, nil
}
