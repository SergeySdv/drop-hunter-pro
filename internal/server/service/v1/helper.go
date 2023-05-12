package v1

import (
	"context"
	"crypto_scripts/internal/defi"
	"crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/server/settings"
	"crypto_scripts/internal/server/user"
	"crypto_scripts/internal/socks5"
	"crypto_scripts/internal/uniclient"
	"strings"

	"github.com/pkg/errors"
)

type HelperService struct {
	v1.UnimplementedHelperServiceServer
	settingsRepository repository.SettingsRepository
	profileRepository  repository.ProfileRepository
}

func NewHelperService(settingsRepository repository.SettingsRepository, profileRepository repository.ProfileRepository) *HelperService {
	return &HelperService{
		settingsRepository: settingsRepository,
		profileRepository:  profileRepository,
	}
}

func (s *HelperService) EstimateStargateBridgeFee(ctx context.Context, req *v1.EstimateStargateBridgeFeeRequest) (*v1.EstimateStargateBridgeFeeResponse, error) {

	profile, err := s.profileRepository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		e := "error estimating fee"
		return &v1.EstimateStargateBridgeFeeResponse{
			Error: &e,
		}, nil
	}
	wallet, err := defi.NewWalletTransactor(profile.MmskPk)
	if err != nil {
		e := "error estimating fee"
		return &v1.EstimateStargateBridgeFeeResponse{
			Error: &e,
		}, nil
	}
	stgs, err := settings.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: req.From,
		UserId:  user.GetUserId(ctx),
		Rep:     s.settingsRepository,
	})
	if err != nil {
		e := "error estimating fee"
		return &v1.EstimateStargateBridgeFeeResponse{
			Error: &e,
		}, nil
	}
	swapper, err := uniclient.NewStargateSwapper(req.From, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: stgs.RpcEndpoint,
	})
	if err != nil {
		e := "error estimating fee"
		return &v1.EstimateStargateBridgeFeeResponse{
			Error: &e,
		}, nil
	}
	fee, err := swapper.GetStargateBridgeFee(ctx, &defi.GetStargateBridgeFeeReq{
		ToChain: req.To,
		Wallet:  wallet.WalletAddr,
	})
	if err != nil {
		e := "error estimating fee"
		return &v1.EstimateStargateBridgeFeeResponse{
			Error: &e,
		}, nil
	}
	return &v1.EstimateStargateBridgeFeeResponse{
		Wei:   fee.Fee1.Int64(),
		Usd:   defi.WEIToUSD(fee.Fee1).String(),
		Eth:   defi.WEIToEther(fee.Fee1).String(),
		Error: nil,
	}, nil

}

func (s *HelperService) ValidatePK(ctx context.Context, req *v1.ValidatePKRequest) (*v1.ValidatePKResponse, error) {
	w, err := defi.NewWalletTransactor(req.Pk)
	if err != nil {
		return &v1.ValidatePKResponse{
			Valid:    false,
			WalletId: nil,
		}, nil
	}
	addr := w.WalletAddr.String()
	return &v1.ValidatePKResponse{
		Valid:    true,
		WalletId: &addr,
	}, nil
}

func (s *HelperService) ValidateProxy(ctx context.Context, req *v1.ValidateProxyRequest) (*v1.ValidateProxyResponse, error) {

	req.Proxy = strings.TrimSpace(req.Proxy)

	if req.Proxy == "" {
		return &v1.ValidateProxyResponse{
			Valid:       true,
			CountryName: "disabled",
			CountryCode: "",
			Ip:          "",
		}, nil
	}
	p, err := socks5.NewSock5ProxyString(req.Proxy)
	if err != nil {
		errMsg := ""
		if errors.Is(err, socks5.ErrParseError) {
			errMsg = "invalid proxy format. want <ip:port:login:password>"
		} else {
			errMsg = "proxy does not responding"
		}

		return &v1.ValidateProxyResponse{
			Valid:        false,
			ErrorMessage: errMsg,
		}, nil
	}

	return &v1.ValidateProxyResponse{
		Valid:       true,
		CountryName: p.Stat.CountryName,
		CountryCode: p.Stat.CountryCode2,
		Ip:          p.Stat.Ip,
	}, nil
}
