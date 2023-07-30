package v1

import (
	"context"

	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/orbiter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type OrbiterService struct {
	v1.UnimplementedOrbiterServiceServer
	s *orbiter.Service
}

func NewOrbiterService(s *orbiter.Service) *OrbiterService {
	return &OrbiterService{
		s: s,
	}
}

func (s *OrbiterService) GetFromTokens(ctx context.Context, req *v1.GetFromTokensReq) (*v1.GetFromTokensRes, error) {
	return &v1.GetFromTokensRes{
		Tokens: s.s.GetFromTokens(req.FromNetwork, req.ToNetwork),
	}, nil
}
func (s *OrbiterService) GetToTokens(ctx context.Context, req *v1.GetToTokensReq) (*v1.GetToTokensRes, error) {
	return &v1.GetToTokensRes{
		Tokens: s.s.GetToTokens(req.FromNetwork, req.ToNetwork, req.FromToken),
	}, nil
}
func (s *OrbiterService) GetSwapOptions(ctx context.Context, req *v1.GetSwapOptionsReq) (*v1.GetSwapOptionsRes, error) {

	opt, ok := s.s.SwapOptions(req.FromNetwork, req.ToNetwork, req.FromToken, req.ToToken)
	if !ok {
		e := "unsupported"
		return &v1.GetSwapOptionsRes{
			Min:   "",
			Max:   "",
			Fee:   "",
			Error: &e,
		}, nil
	}

	return &v1.GetSwapOptionsRes{
		Min:   lib.FloatToString(opt.MinPrice),
		Max:   lib.FloatToString(opt.MaxPrice),
		Fee:   string(opt.GasFee),
		Error: nil,
	}, nil
}
func (s *OrbiterService) GetNetworks(ctx context.Context, req *v1.GetNetworksReq) (*v1.GetNetworksRes, error) {
	return &v1.GetNetworksRes{Networks: s.s.GetNetworks()}, nil
}
