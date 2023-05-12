package uniclient

import (
	"crypto_scripts/internal/exchange"
	"crypto_scripts/internal/exchange/binance"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/socks5"

	"github.com/pkg/errors"
)

func NewExchangeWithdrawer(withdrawer *v1.Withdrawer) (exchange.ExchangeWithdrawer, error) {
	if withdrawer == nil {
		return nil, errors.New("withdrawer is not set")
	}
	proxyString := ""
	if withdrawer.Proxy != nil {
		proxyString = withdrawer.GetProxy()
	}
	proxy, err := socks5.NewSock5ProxyString(proxyString)
	if err != nil {
		return nil, err
	}

	switch withdrawer.ExchangeType {
	case v1.ExchangeType_Binance:
		return binance.NewAdapter(binance.NewService(withdrawer.ApiKey, withdrawer.SecretKey, proxy.Cli)), nil
	default:
		return nil, errors.New("unsupported exchange type: " + withdrawer.ExchangeType.String())
	}
}
