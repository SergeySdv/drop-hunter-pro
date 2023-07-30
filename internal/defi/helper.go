package defi

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

const ETHPrice = 2000

func TokenAmountFloatToWEI(a float64, token Token) *big.Int {

	switch token {
	case v1.Token_USDT, v1.Token_USDC:
		am := big.NewFloat(a)
		am = big.NewFloat(0).Mul(am, big.NewFloat(params.Ether*1e-12))

		r, _ := am.Int(nil)
		return r
	default:
		aa := big.NewFloat(a)
		am := big.NewFloat(0).Mul(aa, big.NewFloat(params.Ether))

		r, _ := am.Int(nil)
		return r
	}
}

func AmountUni(wei *big.Int, network v1.Network) *v1.AmUni {

	gasTokenPrice := pub.Price().ETH
	switch network {
	case v1.Network_ZKSYNCERA, v1.Network_ZKSYNCLITE,
		v1.Network_Etherium, v1.Network_ARBITRUM,
		v1.Network_OPTIMISM, v1.Network_GOERLIETH:
		gasTokenPrice = pub.Price().ETH
	case v1.Network_POLIGON:
		gasTokenPrice = pub.Price().MATIC
	case v1.Network_BinanaceBNB:
		gasTokenPrice = pub.Price().BNB
	case v1.Network_AVALANCHE:
		gasTokenPrice = pub.Price().AVAX
	}

	amEth := WEIToEther(wei)
	amUsd := EthToUsd(amEth, gasTokenPrice)

	weif := big.NewFloat(float64(wei.Int64()))
	gweif := big.NewFloat(0).Quo(weif, big.NewFloat(params.GWei))

	return &v1.AmUni{
		Gwei:    gweif.String(),
		Eth:     amEth.String(),
		Usd:     amUsd.String(),
		Network: network,
		Wei:     wei.String(),
	}
}

func EthToUsd(e *big.Float, price float64) *big.Float {
	return new(big.Float).Mul(e, new(big.Float).SetFloat64(price))
}

func WEIToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func CastFloatToEtherWEI(wei float64) *big.Int {
	f := new(big.Float).Mul(big.NewFloat(wei), big.NewFloat(params.Ether))
	i, _ := f.Int64()

	return big.NewInt(i)
}

func WEIToUSD(wei *big.Int) *big.Float {
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	return new(big.Float).Mul(eth, big.NewFloat(ETHPrice))
}

func WeiToToken(wei *big.Int, token Token) *big.Float {
	switch token {
	case v1.Token_USDT, v1.Token_USDC:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether*1e-12))
	default:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	}
}

func ResolveAmount(amount *v1.Amount, balance *big.Int) (*big.Int, error) {

	if amount == nil {
		return balance, nil
	}

	am := new(big.Int)
	switch amount.Kind.(type) {
	case *v1.Amount_SendAll:
		am = balance
	case *v1.Amount_SendPercent:
		percent := amount.Kind.(*v1.Amount_SendPercent)
		f := math.Round(float64(percent.SendPercent))
		b1p := new(big.Int).Div(balance, new(big.Int).SetInt64(100))
		am.Mul(b1p, new(big.Int).SetInt64(int64(f)))
	case *v1.Amount_SendAmount:
		amountKind := amount.Kind.(*v1.Amount_SendAmount)
		f := new(big.Float).SetFloat64(float64(amountKind.SendAmount))
		am, _ = f.Int(nil)
		//case *v1.Amount_SendValue:
		//	value := amount.Kind.(*v1.Amount_SendValue)
		//	f, err := lib.StringToFloat(value.SendValue)
		//	if err != nil {
		//		return nil, err
		//	}
		//	return TokenAmountFloatToWEI(f, token), nil
	}

	return am, nil
}

func Percent(value *big.Int, percent int) *big.Int {
	f := math.Round(float64(percent))
	b1p := new(big.Int).Div(value, new(big.Int).SetInt64(100))
	return new(big.Int).Mul(b1p, new(big.Int).SetInt64(int64(f)))
}
