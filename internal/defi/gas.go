package defi

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var ErrUserGasLimitToLow = errors.New("gas is higher than user gas limit settings")

func ErrGasHigh(want, have *big.Int) error {
	return errors.Wrap(ErrUserGasLimitToLow, " want:"+want.String()+" have: "+have.String())
}

func (c *EtheriumClient) ResolveGasPrice(limit *GasLimit, tx *types.Transaction) *big.Int {
	if limit == nil {
		limit = c.c.GasLimit
	}

	if limit == nil {
		return nil
	}

	println("limit.TotalGas", limit.TotalGas.String())
	println("network.gasPrice", tx.GasPrice().String())
	if limit.TotalGas != nil && limit.TotalGas.Cmp(big.NewInt(0)) == 1 {

		gasLimitPrice := big.NewInt(int64(tx.Gas()))
		gasPrice := tx.GasPrice()
		totalPrice := big.NewInt(0).Mul(gasLimitPrice, gasPrice)

		println("network.total", totalPrice.String())

		// цена газа выше заданной
		if totalPrice.Cmp(limit.TotalGas) == 1 {
			return big.NewInt(0).Quo(limit.TotalGas, gasLimitPrice)
		}
	}

	return tx.GasPrice()
}
