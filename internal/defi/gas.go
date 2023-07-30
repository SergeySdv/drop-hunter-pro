package defi

import (
	"math/big"

	"github.com/pkg/errors"
)

var ErrUserGasLimitToLow = errors.New("gas is higher than user gas limit settings")

func ErrGasHigh(want, have *big.Int) error {
	return errors.Wrap(ErrUserGasLimitToLow, " want:"+want.String()+" have: "+have.String())
}

func MinerGasLegacy(gasPrice *big.Int, gasLimit uint64) *big.Int {
	return big.NewInt(0).Mul(gasPrice, big.NewInt(0).SetUint64(gasLimit))
}

func MinerGas(baseFee, priorityFee *big.Int, gasLimit uint64) *big.Int {

	gasPrice := new(big.Int).Add(baseFee, priorityFee)

	return new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))
}

func ResolveGasPriceZksync(gasMax, gas, price *big.Int) *big.Int {
	if gasMax == nil {
		gasMax = big.NewInt(0)
	}
	if gasMax.Cmp(big.NewInt(0)) == 0 {
		return price
	}
	estimatedGas := new(big.Int).Mul(gas, price)

	if estimatedGas.Cmp(gasMax) > 0 {
		return new(big.Int).Div(gasMax, price)
	}

	return price
}

func ResolveGasPriceLayerZero(gasMax, gas, price *big.Int) *big.Int {
	if gasMax == nil {
		gasMax = big.NewInt(0)
	}
	if gasMax.Cmp(big.NewInt(0)) == 0 {
		return price
	}
	estimatedGas := new(big.Int).Mul(gas, price)

	if estimatedGas.Cmp(gasMax) > 0 {
		return new(big.Int).Div(gasMax, price)
	}

	return price
}
