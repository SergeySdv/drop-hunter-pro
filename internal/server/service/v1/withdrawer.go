package v1

import (
	"context"
	"crypto_scripts/internal/exchange/binance"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/socks5"
)

type WithdrawerService struct {
	v1.UnimplementedWithdrawerServiceServer
	repository repository.WithdrawerRepository
}

func NewWithdrawerService(repository repository.WithdrawerRepository) *WithdrawerService {
	return &WithdrawerService{
		repository: repository,
	}
}

func (s *WithdrawerService) CreateWithdrawer(ctx context.Context, req *v1.CreateWithdrawerRequest) (*v1.CreateWithdrawerResponse, error) {
	return s.repository.CreateWithdrawer(ctx, req)
}
func (s *WithdrawerService) ListWithdrawer(ctx context.Context, req *v1.ListWithdrawerRequest) (*v1.ListWithdrawerResponse, error) {
	return s.repository.ListWithdrawers(ctx, req)
}
func (s *WithdrawerService) DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error) {
	if _, err := s.repository.DeleteWithdrawer(ctx, req); err != nil {
		return nil, err
	}
	return &v1.DeleteWithdrawerResponse{}, nil
}

func (s *WithdrawerService) GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error) {
	w, err := s.repository.GetWithdrawers(ctx, req.WithdrawerId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if w.Proxy.Valid {
		proxyString = w.Proxy.String
	}
	proxy, err := socks5.NewSock5ProxyString(proxyString)
	if err != nil {
		return nil, err
	}

	a := binance.NewAdapter(binance.NewService(w.ApiKey, w.SecretKey, proxy.Cli))

	return a.GetExchangeWithdrawOptions(ctx, req)
}
