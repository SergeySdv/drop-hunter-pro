package task

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

func TestResolveAmount(t *testing.T) {
	type args struct {
		amount  *v1.Amount
		balance *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "all",
			args: args{
				amount: &v1.Amount{
					Kind: &v1.Amount_SendAll{
						SendAll: true,
					},
				},
				balance: new(big.Int).SetInt64(100),
			},
			want: new(big.Int).SetInt64(100),
		},
		{
			name: "percent",
			args: args{
				amount: &v1.Amount{
					Kind: &v1.Amount_SendPercent{
						SendPercent: 10.23,
					},
				},
				balance: new(big.Int).SetInt64(100),
			},
			want: new(big.Int).SetInt64(10),
		},
		{
			name: "percent",
			args: args{
				amount: &v1.Amount{
					Kind: &v1.Amount_SendPercent{
						SendPercent: 1,
					},
				},
				balance: new(big.Int).SetInt64(100),
			},
			want: new(big.Int).SetInt64(1),
		},

		{
			name: "percent",
			args: args{
				amount: &v1.Amount{
					Kind: &v1.Amount_SendPercent{
						SendPercent: 24,
					},
				},
				balance: new(big.Int).SetInt64(200),
			},
			want: new(big.Int).SetInt64(48),
		},
		{
			name: "am",
			args: args{
				amount: &v1.Amount{
					Kind: &v1.Amount_SendAmount{
						SendAmount: 10,
					},
				},
				balance: new(big.Int).SetInt64(100),
			},
			want: new(big.Int).SetInt64(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := defi.ResolveAmount(tt.args.amount, tt.args.balance, v1.Token_ETH); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
