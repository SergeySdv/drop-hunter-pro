package defi

import (
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

const ETHPrice = 2000

func TokenAmountToWEI(a int64, token Token) *big.Int {
	switch token {
	case v1.Token_USDT, v1.Token_USDC:
		am := big.NewInt(a)
		am = big.NewInt(0).Mul(am, big.NewInt(params.Ether*1e-12))
		return am
	default:
		am := big.NewInt(a)
		am = big.NewInt(0).Mul(am, big.NewInt(params.Ether))
		return am
	}

}

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

func WEIToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
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
