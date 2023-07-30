package defi

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	type args struct {
		a *RollArg
	}
	tests := []struct {
		name string
		args args
		want RollRet
	}{
		{
			name: "balance >>> Amount",
			args: args{
				a: &RollArg{
					MinerGas: big.NewInt(1),
					GasExt:   big.NewInt(1),
					Balance:  big.NewInt(10),
					Amount:   big.NewInt(2),
				},
			},
			want: RollRet{
				Value:  big.NewInt(3),
				Amount: big.NewInt(2),
			},
		},
		{
			name: "balance <<< Amount",
			args: args{
				a: &RollArg{
					MinerGas: big.NewInt(1),
					GasExt:   big.NewInt(1),
					Balance:  big.NewInt(10),
					Amount:   big.NewInt(20),
				},
			},
			want: RollRet{
				Value:  big.NewInt(9),
				Amount: big.NewInt(8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Roll(tt.args.a), "Roll(%v)", tt.args.a)
		})
	}
}

func TestEthToUsd(t *testing.T) {
	type args struct {
		e     *big.Float
		price float64
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{
			name: "",
			args: args{
				e:     big.NewFloat(0.001),
				price: 2000,
			},
			want: big.NewFloat(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EthToUsd(tt.args.e, tt.args.price), "EthToUsd(%v, %v)", tt.args.e, tt.args.price)
		})
	}
}
