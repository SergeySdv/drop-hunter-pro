package uniclient

import (
	"context"
	"strings"

	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/exchange/binance"
	"github.com/hardstylez72/cry/internal/exchange/okex"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/pkg/errors"
)

func NewExchangeWithdrawer(withdrawer *repository.Withdrawer, profile *repository.Profile) (exchange.Withdrawer, error) {
	if profile == nil {
		profile = &repository.Profile{
			UserAgent: lib.DefaultUserAgent,
		}
	}

	if withdrawer == nil {
		return nil, errors.New("withdrawer is not set")
	}
	proxyString := ""
	if withdrawer.Proxy.Valid {
		proxyString = withdrawer.Proxy.String
	}

	proxy, err := socks5.NewSock5ProxyString(proxyString, profile.UserAgent)
	if err != nil {
		return nil, err
	}

	switch withdrawer.ExchangeType {
	case v1.ExchangeType_Binance.String():
		return binance.NewAdapter(binance.NewService(string(withdrawer.ApiKey), string(withdrawer.SecretKey), proxy.Cli)), nil
	case v1.ExchangeType_Okex.String():

		subs := strings.Split(string(withdrawer.ApiKey), ":")
		if len(subs) != 2 {
			return nil, errors.New("okex withdrawer.ApiKey is invalid")
		}
		return okex.NewService(context.Background(), &okex.Config{
			ApiKey:     subs[0],
			SecretKey:  string(withdrawer.SecretKey),
			PassPhrase: subs[1],
			HttpClient: proxy.Cli,
		})
	default:
		return nil, errors.New("unsupported exchange type: " + withdrawer.ExchangeType)
	}
}
