package v1

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/exchange/binance"
	"github.com/hardstylez72/cry/internal/exchange/okex"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WithdrawerService struct {
	v1.UnimplementedWithdrawerServiceServer
	repository        repository.WithdrawerRepository
	userRepository    repository.UserRepository
	profileRepository repository.ProfileRepository
}

func NewWithdrawerService(repository repository.WithdrawerRepository, userRepository repository.UserRepository, profileRepository repository.ProfileRepository) *WithdrawerService {
	return &WithdrawerService{
		repository:        repository,
		userRepository:    userRepository,
		profileRepository: profileRepository,
	}
}

func (s *WithdrawerService) GetExchangeDepositOptions(ctx context.Context, req *v1.GetExchangeDepositOptionsRequest) (*v1.GetExchangeDepositOptionsResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	wd, err := s.repository.GetWithdrawer(ctx, req.WithdrawerId, userId)
	if err != nil {
		return nil, err
	}

	userAgent := lib.DefaultUserAgent

	proxy, err := socks5.NewSock5ProxyString(wd.Proxy.String, userAgent)
	if err != nil {
		return nil, err
	}

	binanceCli := binance.NewService(string(wd.ApiKey), string(wd.SecretKey), proxy.Cli)

	addr, err := binanceCli.GetDepositAddr(ctx, req.Network, req.Token)
	if err != nil {
		return nil, err
	}

	return &v1.GetExchangeDepositOptionsResponse{
		Addr: addr,
	}, nil
}
func (s *WithdrawerService) OkexDepositAddrAttach(ctx context.Context, req *v1.OkexDepositAddrAttachRequest) (*v1.OkexDepositAddrAttachResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	err = s.repository.OkexDepositAddrAttach(ctx, &repository.OkexDepositAddr{
		WithdrawerId: req.WithdrawerId,
		Addr:         req.OkexDepositAddr,
		ProfileId:    req.ProfileId,
		UserId:       userId,
	})

	if err != nil {
		return nil, err
	}
	return &v1.OkexDepositAddrAttachResponse{}, nil
}
func (s *WithdrawerService) OkexDepositAddrDetach(ctx context.Context, req *v1.OkexDepositAddrDetachRequest) (*v1.OkexDepositAddrDetachResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	err = s.repository.OkexDepositAddrDetach(ctx, &repository.OkexDepositAddr{
		WithdrawerId: req.WithdrawerId,
		Addr:         req.OkexDepositAddr,
		ProfileId:    req.ProfileId,
		UserId:       userId,
	})

	if err != nil {
		return nil, err
	}
	return &v1.OkexDepositAddrDetachResponse{}, nil
}
func (s *WithdrawerService) ListDepositAddresses(ctx context.Context, req *v1.ListDepositAddressesRequest) (*v1.ListDepositAddressesResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	wd, err := s.repository.GetWithdrawer(ctx, req.WithdrawerId, userId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if wd.Proxy.Valid {
		proxyString = wd.Proxy.String
	}

	userAgent := lib.DefaultUserAgent

	proxy, err := socks5.NewSock5ProxyString(proxyString, userAgent)
	if err != nil {
		return nil, err
	}

	sub := strings.Split(string(wd.ApiKey), ":")

	cli, err := okex.NewService(ctx, &okex.Config{
		ApiKey:     sub[0],
		SecretKey:  string(wd.SecretKey),
		PassPhrase: sub[1],
		HttpClient: proxy.Cli,
	})
	if err != nil {
		return nil, err
	}

	attached, err := s.repository.ListDepositAddrAttach(ctx, userId)
	if err != nil {
		return nil, err
	}
	items := make([]*v1.DepositAddresses, 0)
	addresses, err := cli.GetDepositAddress(ctx, "USDT")
	if err != nil {
		return nil, err
	}

	for _, v := range addresses {

		var ProfileId string
		var ProfileLabel string
		_, exist := attached[v.Addr]
		if exist {
			ProfileId = attached[v.Addr].ProfileId
			ProfileLabel = attached[v.Addr].ProfileLabel
		}

		items = append(items, &v1.DepositAddresses{
			Addr:         v.Addr,
			Tag:          &v.Tag,
			ProfileId:    &ProfileId,
			ProfileLabel: &ProfileLabel,
		})
	}

	return &v1.ListDepositAddressesResponse{
		Items: items,
	}, nil
}
func (s *WithdrawerService) ListSubWithdrawer(ctx context.Context, req *v1.ListSubWithdrawerRequest) (*v1.ListSubWithdrawerResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	list, err := s.repository.ListSubWithdrawers(ctx, req.WithdrawerId, userId)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Withdrawer, 0)

	for _, l := range list {
		out = append(out, l.ToPB())
	}

	return &v1.ListSubWithdrawerResponse{Withdrawers: out}, nil
}
func (s *WithdrawerService) CreateSubWithdrawer(ctx context.Context, req *v1.CreateOkexWithdrawerRequest) (*v1.CreateOkexWithdrawerResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	a := &repository.Withdrawer{
		Id:           uuid.New().String(),
		ExchangeType: v1.ExchangeType_Okex.String(), //todo: заменть
		Label:        req.Label,
		Proxy: sql.NullString{
			String: req.GetProxy(),
			Valid:  true,
		},
		SecretKey: []byte(req.SecretKey),
		ApiKey:    []byte(req.ApiKey),
		UserId:    userId,
		CreatedAt: time.Now(),
		PrevId: sql.NullString{
			String: req.ParentId,
			Valid:  true,
		},
	}

	wd, err := uniclient.NewExchangeWithdrawer(a, nil)
	if err != nil {
		e := err.Error()
		return &v1.CreateOkexWithdrawerResponse{
			Error: &e,
		}, nil
	}

	if err := wd.Ping(ctx); err != nil {
		e := err.Error()
		return &v1.CreateOkexWithdrawerResponse{
			Error: &e,
		}, nil
	}

	if err = s.repository.CreateSubWithdrawer(ctx, a); err != nil {
		return nil, errors.Wrap(err, "repository.CreateSubWithdrawer")
	}

	return &v1.CreateOkexWithdrawerResponse{}, nil
}
func (s *WithdrawerService) CreateWithdrawer(ctx context.Context, req *v1.CreateWithdrawerRequest) (*v1.CreateWithdrawerResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	a := &repository.Withdrawer{
		Id:           uuid.New().String(),
		ExchangeType: req.ExchangeType.String(),
		Label:        req.Label,
		Proxy: sql.NullString{
			String: req.GetProxy(),
			Valid:  true,
		},
		SecretKey: []byte(req.SecretKey),
		ApiKey:    []byte(req.ApiKey),
		UserId:    userId,
		CreatedAt: time.Time{},
	}

	wd, err := uniclient.NewExchangeWithdrawer(a, nil)
	if err != nil {
		return nil, err
	}

	if err := wd.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "can't connect to exchange")
	}

	res, err := s.repository.CreateWithdrawer(ctx, a)
	if err != nil {
		return nil, err
	}
	return &v1.CreateWithdrawerResponse{
		Withdrawer: res.ToPB(),
	}, nil
}
func (s *WithdrawerService) ListWithdrawer(ctx context.Context, req *v1.ListWithdrawerRequest) (*v1.ListWithdrawerResponse, error) {

	userid, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	list, err := s.repository.ListWithdrawers(ctx, userid)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Withdrawer, 0)

	for _, l := range list {
		out = append(out, l.ToPB())
	}

	return &v1.ListWithdrawerResponse{Withdrawers: out}, nil
}
func (s *WithdrawerService) DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error) {
	if _, err := s.repository.DeleteWithdrawer(ctx, req); err != nil {
		return nil, err
	}
	return &v1.DeleteWithdrawerResponse{}, nil
}
func (s *WithdrawerService) GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	w, err := s.repository.GetWithdrawer(ctx, req.WithdrawerId, userId)
	if err != nil {
		return nil, err
	}

	a, err := uniclient.NewExchangeWithdrawer(w, nil)
	if err != nil {
		return nil, err
	}

	return a.GetExchangeWithdrawOptions(ctx, req)
}
func (s *WithdrawerService) UpdateWithdrawer(ctx context.Context, req *v1.UpdateWithdrawerRequest) (*v1.UpdateWithdrawerResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	source, err := s.repository.GetWithdrawer(ctx, req.WithdrawerId, userId)
	if err != nil {
		return nil, err
	}

	source.Label = req.Label
	source.Proxy.String = req.Proxy

	wd, err := uniclient.NewExchangeWithdrawer(source, nil)
	if err != nil {
		e := err.Error()
		return &v1.UpdateWithdrawerResponse{
			Error: &e,
		}, err
	}

	if err := wd.Ping(ctx); err != nil {
		e := err.Error()
		return &v1.UpdateWithdrawerResponse{
			Error: &e,
		}, err
	}

	err = s.repository.UpdateWithdrawer(ctx, &repository.Withdrawer{
		Id:    req.WithdrawerId,
		Label: req.Label,
		Proxy: sql.NullString{
			String: req.Proxy,
			Valid:  true,
		},
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	res := &v1.UpdateWithdrawerResponse{}
	return res, nil
}
func (s *WithdrawerService) GetWithdrawer(ctx context.Context, req *v1.GetWithdrawerRequest) (*v1.GetWithdrawerResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	source, err := s.repository.GetWithdrawer(ctx, req.GetWithdrawerId(), userId)
	if err != nil {
		return nil, err
	}

	w := source

	wd, err := uniclient.NewExchangeWithdrawer(w, nil)
	if err != nil {
		e := err.Error()
		return &v1.GetWithdrawerResponse{
			Withdrawer: w.ToPB(),
			Error:      &e,
		}, nil
	}

	if err := wd.Ping(ctx); err != nil {
		e := err.Error()
		return &v1.GetWithdrawerResponse{
			Withdrawer: w.ToPB(),
			Error:      &e,
		}, nil
	}

	return &v1.GetWithdrawerResponse{
		Withdrawer: w.ToPB(),
	}, nil
}
func (s *WithdrawerService) Withdraw(ctx context.Context, req *v1.WithdrawReq) (*v1.WithdrawRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	source, err := s.repository.GetWithdrawer(ctx, req.GetWithdrawerId(), userId)
	if err != nil {
		return nil, err
	}

	w := source

	wd, err := uniclient.NewExchangeWithdrawer(w, nil)
	if err != nil {
		e := err.Error()
		return &v1.WithdrawRes{
			ErrorMessage: &e,
		}, nil
	}

	if err := wd.Ping(ctx); err != nil {
		e := err.Error()
		return &v1.WithdrawRes{
			ErrorMessage: &e,
		}, nil
	}

	profile, err := s.profileRepository.GetProfile(ctx, req.GetProfileId())
	if err != nil {
		e := err.Error()
		return &v1.WithdrawRes{
			ErrorMessage: &e,
		}, nil
	}

	wallet, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		e := err.Error()
		return &v1.WithdrawRes{
			ErrorMessage: &e,
		}, nil
	}

	res, err := wd.Withdraw(ctx, &exchange.WithdrawRequest{
		ToAddress: wallet.WalletAddrHR,
		Amount:    req.Amount,
		Network:   req.Network,
		Token:     req.Token,
	})
	if err != nil {
		e := err.Error()
		return &v1.WithdrawRes{
			ErrorMessage: &e,
		}, nil
	}

	return &v1.WithdrawRes{
		WithdrawId: res.WithdrawId,
	}, nil
}
func (s *WithdrawerService) WithdrawStatus(ctx context.Context, req *v1.WithdrawStatusReq) (*v1.WithdrawStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawStatus not implemented")
}

//func (s *WithdrawerService) ExportExchangeAccounts(ctx context.Context, req *v1.ExportExchangeAccountsReq) (*v1.ExportExchangeAccountsRes, error) {
//	userId, err := user.GetUserId(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	accounts, err := s.repository.ExportExchangeAccounts(ctx, userId)
//	if err != nil {
//		return nil, err
//	}
//
//	out := make([]AccountExport, 0)
//	sub := map[string][]AccountExport{}
//
//	//sub accs
//	for i := range accounts {
//		a := &accounts[i]
//		if !a.PrevId.Valid {
//			continue
//		}
//		sub[a.PrevId.String] = append(sub[a.PrevId.String], ParseExportAcc(a, nil))
//	}
//	for i := range accounts {
//		a := &accounts[i]
//		if a.PrevId.Valid {
//			continue
//		}
//		out = append(out, ParseExportAcc(a, sub[a.Id]))
//	}
//	marshal, err := json.MarshalIndent(&out, "", "    ")
//	if err != nil {
//		return nil, err
//	}
//
//	return &v1.ExportExchangeAccountsRes{
//		Text: string(marshal),
//	}, nil
//}

func ParseExportAcc(in *repository.Withdrawer, sub []AccountExport) AccountExport {
	passphrase := ""
	apiKey := string(in.ApiKey)
	if in.ExchangeType == v1.ExchangeType_Okex.String() {
		sub := strings.Split(string(in.ApiKey), ":")
		passphrase = sub[1]
		apiKey = sub[0]
	}

	return AccountExport{
		ExchangeType: in.ExchangeType,
		Label:        in.Label,
		Proxy:        in.Proxy.String,
		SecretKey:    string(in.SecretKey),
		ApiKey:       apiKey,
		PassPhrase:   passphrase,
		SubAccs:      sub,
	}
}

type AccountExport struct {
	ExchangeType string          `json:"exchange_type"`
	Label        string          `json:"label"`
	Proxy        string          `json:"proxy"`
	SecretKey    string          `json:"secret_key"`
	ApiKey       string          `json:"api_key"`
	PassPhrase   string          `json:"pass_phrase"`
	SubAccs      []AccountExport `json:"sub_accs"`
}
