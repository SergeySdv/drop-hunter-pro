package v1

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"sync"
	"time"

	"github.com/corpix/uarand"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type ProfileService struct {
	v1.UnimplementedProfileServiceServer
	repository      repository.ProfileRepository
	settingsService *settings.Service
}

func NewProfileService(repository repository.ProfileRepository, settingsService *settings.Service) *ProfileService {

	return &ProfileService{
		repository:      repository,
		settingsService: settingsService,
	}
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (*v1.UpdateProfileResponse, error) {

	p := &repository.Profile{
		Id:        req.ProfileId,
		Label:     req.Label,
		Proxy:     sql.NullString{},
		Meta:      sql.NullString{},
		UserAgent: req.UserAgent,
	}

	if req.Meta != nil {
		p.Meta.Valid = true
		p.Meta.String = req.GetMeta()
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
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	valid, err := s.repository.ValidateLabel(ctx, &repository.ValidateLabelReq{
		ProfileId: req.ProfileId,
		Label:     req.Label,
		UserId:    userId,
	})
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

	pr, err := p.ToPB()
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileResponse{
		Profile: pr,
	}, nil
}
func (s *ProfileService) CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	a := &repository.Profile{
		Id:        uuid.New().String(),
		Label:     req.Label,
		Proxy:     sql.NullString{},
		MmskPk:    []byte(req.MmskPk),
		Meta:      sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
		UserAgent: uarand.NewWithCustomList(lib.UserAgents).GetRandom(),
	}

	w, err := defi.NewWalletTransactor(req.MmskPk)
	if err != nil {
		return nil, errors.New("invalid pk")
	}
	a.MmskId = []byte(w.WalletAddrHR)

	if req.Proxy != nil {
		p, err := socks5.NewSock5ProxyString(req.GetProxy(), "")
		if err != nil {
			return nil, errors.New("invalid proxy")
		}
		_, err = p.GetIpStat(ctx)
		if err != nil {
			return nil, errors.New("proxy stat is not available")
		}
	}

	if req.Meta != nil {
		a.Meta.Valid = true
		a.Meta.String = *req.Meta
	}

	if req.Proxy != nil {
		a.Proxy.Valid = true
		a.Proxy.String = *req.Proxy
	}

	err = s.repository.CreateProfile(ctx, a)
	if err != nil {
		if errors.Is(err, pg.EntityAlreadyExist) {
			return nil, errors.New("pk already exist")
		}
		return nil, err
	}

	res, err := s.repository.GetProfile(ctx, a.Id)
	if err != nil {
		return nil, err
	}
	pr, err := res.ToPB()
	if err != nil {
		return nil, err
	}

	return &v1.CreateProfileResponse{
		Profile: pr,
	}, nil
}
func (s *ProfileService) ListProfile(ctx context.Context, req *v1.ListProfileRequest) (*v1.ListProfileResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.repository.ListProfiles(ctx, userId)
	if err != nil {
		return nil, err
	}

	profiles := make([]*v1.Profile, 0)

	for i := range res {
		p := res[i]
		pp, err := p.ToPB()
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, pp)
	}

	return &v1.ListProfileResponse{
		Profiles: profiles,
	}, nil
}
func (s *ProfileService) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {
	_, _ = s.repository.DeleteProfile(ctx, req)
	return &v1.DeleteProfileResponse{}, nil
}
func (s *ProfileService) SearchProfilesNotConnectedToOkexDeposit(ctx context.Context, _ *v1.SearchProfilesNotConnectedToOkexDepositRequest) (*v1.SearchProfilesNotConnectedToOkexDepositResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	profiles, err := s.repository.SearchNotConnectedOkexDepositProfile(ctx, userId)
	if err != nil {
		return nil, err
	}

	out := make([]*v1.Profile, 0)

	for _, p := range profiles {
		o, err := p.ToPB()
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return &v1.SearchProfilesNotConnectedToOkexDepositResponse{Profiles: out}, nil
}
func (s *ProfileService) SearchProfile(ctx context.Context, req *v1.SearchProfileRequest) (*v1.SearchProfileResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	profiles, err := s.repository.SearchProfile(ctx, req.Pattern, userId)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Profile, 0)

	for _, p := range profiles {
		o, err := p.ToPB()
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return &v1.SearchProfileResponse{Profiles: out}, nil
}
func (s *ProfileService) GetBalance(ctx context.Context, req *v1.GetBalanceRequest) (*v1.GetBalanceResponse, error) {

	tokens := []v1.Token{
		v1.Token_USDT,
		v1.Token_USDC,
		v1.Token_ETH,
		v1.Token_STG,
		v1.Token_WETH,
		v1.Token_LSD,
		v1.Token_LUSD,
		v1.Token_MUTE,
		v1.Token_MAV,
		v1.Token_SPACE,
		v1.Token_VC,
		v1.Token_IZI,
	}

	var err error
	wg := sync.WaitGroup{}

	profile, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	wallet, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		return nil, err
	}

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	stgs, err := s.settingsService.GetSettingsNetwork(ctx, &settings.GetSettingsNetworkRequest{
		Network: req.Network,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}

	cli, err := uniclient.NewBaseClient(req.Network, &uniclient.BaseClientConfig{
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

		go func(balancer defi.Balancer, i int) {
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

//	func (s *ProfileService) ExportProfiles(ctx context.Context, req *v1.ExportProfilesReq) (*v1.ExportProfilesRes, error) {
//		userId, err := user.GetUserId(ctx)
//		if err != nil {
//			return nil, err
//		}
//
//		profiles, err := s.repository.ExportProfiles(ctx, userId)
//		if err != nil {
//			return nil, err
//		}
//
//		buf := bytes.NewBuffer([]byte{})
//		buf.Grow(100000)
//		w := csv.NewWriter(buf)
//
//		w.UseCRLF = true
//		w.Comma = ';'
//		for i, p := range profiles {
//			if i%10 == 0 {
//				if err := w.Write([]string{}); err != nil {
//					return nil, err
//				}
//			}
//			if err := w.Write([]string{string(p.MmskPk), p.Proxy.String, p.Label}); err != nil {
//				return nil, err
//			}
//		}
//		w.Flush()
//		return &v1.ExportProfilesRes{
//			Text: buf.String(),
//		}, nil
//	}
func (s *ProfileService) GenerateProfiles(ctx context.Context, req *v1.GenerateProfilesReq) (*v1.GenerateProfilesRes, error) {

	buf := bytes.NewBuffer([]byte{})
	buf.Grow(100000)
	w := csv.NewWriter(buf)

	w.UseCRLF = true
	w.Comma = ';'

	for i := 0; i < int(req.Count); i++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		//address := crypto.PubkeyToAddress(key.PublicKey).Hex()

		// Get the private key
		privateKey := hex.EncodeToString(key.D.Bytes())

		if err := w.Write([]string{privateKey, "", ""}); err != nil {
			return nil, err
		}
	}

	w.Flush()

	return &v1.GenerateProfilesRes{
		Text: buf.String(),
	}, nil

}
