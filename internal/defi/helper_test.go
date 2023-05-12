package defi

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

func TestTokenAmountFloatToWEI(t *testing.T) {
	type args struct {
		a float64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "",
			args: args{
				1,
			},
			want: big.NewInt(0).Mul(big.NewInt(1), big.NewInt(params.Ether)),
		},
		{
			name: "",
			args: args{
				1.23456,
			},
			want: big.NewInt(0).Mul(big.NewInt(123456), big.NewInt(params.Ether*10e-5)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenAmountFloatToWEI(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenAmountFloatToWEI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConv(t *testing.T) {
	a := TokenAmountFloatToWEI(29)
	b := TokenAmountToWEI(29)

	assert.True(t, a.Cmp(b) == 0)
}
