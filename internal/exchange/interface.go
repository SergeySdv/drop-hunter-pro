package exchange

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
)

type WithdrawId = string

type WithdrawRequest struct {
	ToAddress string
	Amount    string
	Network   string
	Token     string
}

type WithdrawResponse struct {
	WithdrawId WithdrawId
}

type ExchangeWithdrawer interface {
	GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error)
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
	WaitConfirm(ctx context.Context, id WithdrawId) (*string, error)
}
