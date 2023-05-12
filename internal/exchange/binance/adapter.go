package binance

import (
	"context"
	"crypto_scripts/internal/exchange"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"time"
)

type Adapter struct {
	s *Service
}

func NewAdapter(s *Service) *Adapter {
	return &Adapter{
		s: s,
	}
}

func (a *Adapter) GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error) {

	assets, err := a.s.GetAllUserAssets(ctx)
	if err != nil {
		return nil, err
	}

	options, err := a.s.GetWithDrawOptions(ctx)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.ExchangeWithdrawOptions, 0)

	for _, o := range options {
		networks := make([]*v1.ExchangeWithdrawNetwork, 0)

		amount, tokenHasBalance := assets[o.Token]
		if !tokenHasBalance {
			continue
		}

		for _, n := range o.Networks {
			networks = append(networks, &v1.ExchangeWithdrawNetwork{
				Network:   n.Network,
				MinAmount: n.MinAmount,
				MaxAmount: n.MaxAmount,
				Fee:       n.Fee,
			})
		}

		tmp = append(tmp, &v1.ExchangeWithdrawOptions{
			Token:    o.Token,
			Amount:   amount,
			Networks: networks,
		})
	}

	return &v1.GetExchangeWithdrawOptionsResponse{
		Tokens: tmp,
	}, nil
}

func (a *Adapter) Withdraw(ctx context.Context, req *exchange.WithdrawRequest) (*exchange.WithdrawResponse, error) {
	res, err := a.s.Withdraw(ctx, &WithdrawReq{
		ToAddress: req.ToAddress,
		Amount:    req.Amount,
		Network:   req.Network,
		Coin:      req.Token,
	})
	if err != nil {
		return nil, err
	}

	return &exchange.WithdrawResponse{
		WithdrawId: res.WithDrawOrderId,
	}, nil
}

func (a *Adapter) WaitConfirm(ctx context.Context, id exchange.WithdrawId) (*string, error) {
	return a.s.WithdrawsPending(ctx, id, time.Second*10)
}
