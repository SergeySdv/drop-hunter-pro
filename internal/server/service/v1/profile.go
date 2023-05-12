package v1

import (
	"context"
	"crypto_scripts/internal/defi"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/server/settings"
	"crypto_scripts/internal/server/user"
	"crypto_scripts/internal/uniclient"
	"database/sql"
	"sync"
)

type ProfileService struct {
	v1.UnimplementedProfileServiceServer
	repository         repository.ProfileRepository
	settingsRepository repository.SettingsRepository
}

func NewProfileService(repository repository.ProfileRepository, settingsRepository repository.SettingsRepository) *ProfileService {
	return &ProfileService{
		repository:         repository,
		settingsRepository: settingsRepository,
	}
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (*v1.UpdateProfileResponse, error) {

	p := &repository.Profile{
		Id:              req.ProfileId,
		Label:           req.Label,
		Proxy:           sql.NullString{},
		Meta:            sql.NullString{},
		OkexDepositAddr: sql.NullString{},
		OkexAccName:     sql.NullString{},
	}

	if req.Meta != nil {
		p.Meta.Valid = true
		p.Meta.String = req.GetMeta()
	}

	if req.OkexAccount != nil {
		p.OkexDepositAddr.Valid = true
		p.OkexDepositAddr.String = req.GetOkexAccount().GetDepositAddr()

		p.OkexAccName.Valid = true
		p.OkexAccName.String = req.GetOkexAccount().GetSubAccName()
	}

	if req.Proxy != nil {
		p.Proxy.Valid = true
		p.Proxy.String = req.GetProxy()
	}

	if err := s.repository.UpdateProfile(ctx, p); err != nil {
		return nil, err
	}
	return &v1.UpdateProfileResponse{}, nil
}

func (s *ProfileService) ValidateLabel(ctx context.Context, req *v1.ValidateLabelRequest) (*v1.ValidateLabelResponse, error) {

	valid, err := s.repository.ValidateLabel(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.ValidateLabelResponse{
		Valid: *valid,
	}, nil
}

func (s *ProfileService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.GetProfileResponse, error) {
	p, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponse{
		Profile: p.ToPB(),
	}, nil
}

func (s *ProfileService) CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error) {
	return s.repository.CreateProfile(ctx, req)
}
func (s *ProfileService) ListProfile(ctx context.Context, req *v1.ListProfileRequest) (*v1.ListProfileResponse, error) {
	return s.repository.ListProfiles(ctx, req)
}
func (s *ProfileService) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {
	_, _ = s.repository.DeleteProfile(ctx, req)
	return &v1.DeleteProfileResponse{}, nil
}
func (s *ProfileService) SearchProfile(ctx context.Context, req *v1.SearchProfileRequest) (*v1.SearchProfileResponse, error) {
	return s.repository.SearchProfile(ctx, req)
}

func (s *ProfileService) GetBalance(ctx context.Context, req *v1.GetBalanceRequest) (*v1.GetBalanceResponse, error) {

	tokens := []v1.Token{
		v1.Token_USDT,
		v1.Token_USDC,
		v1.Token_ETH,
		v1.Token_STG,
	}

	var err error
	wg := sync.WaitGroup{}

	profile, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	wallet, err := defi.NewWalletTransactor(profile.MmskPk)
	if err != nil {
		return nil, err
	}

	stgs, err := settings.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: req.Network,
		UserId:  user.GetUserId(ctx),
		Rep:     s.settingsRepository,
	})
	if err != nil {
		return nil, err
	}

	cli, err := uniclient.NewBaseClient(req.Network, &uniclient.BaseClientConfig{
		ProxyString: "",
		RPCEndpoint: stgs.RpcEndpoint,
	})
	if err != nil {
		return nil, err
	}

	tokens = append(tokens, cli.GetNetworkToken())

	m := make(map[v1.Token]bool)
	for _, token := range tokens {
		m[token] = true
	}

	wg.Add(len(m))
	balance := make([]*v1.Balance, len(m))

	for i := range tokens {
		_, ok := m[tokens[i]]
		if ok {
			delete(m, tokens[i])
		} else {
			continue
		}

		go func(balancer defi.Networker, i int) {
			defer wg.Done()

			b, err := balancer.GetBalance(ctx, &defi.GetBalanceReq{
				WalletAddress: wallet.WalletAddr,
				Token:         tokens[i],
			})
			if err != nil {
				msg := err.Error()
				balance[i] = &v1.Balance{Token: tokens[i], Error: &msg}
			} else {
				balance[i] = &v1.Balance{Token: tokens[i], Amount: b.HumanReadable, Wei: b.WEI.String()}
			}

		}(cli, i)
	}

	wg.Wait()

	tmp := make([]*v1.Balance, 0)
	for _, b := range balance {

		if b.Error != nil {
			continue
		}

		if b == nil {
			continue
		}
		if b.Wei == "0" {
			continue
		}
		tmp = append(tmp, b)
	}

	return &v1.GetBalanceResponse{
		Balance: tmp,
	}, nil
}
