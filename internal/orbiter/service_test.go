package orbiter

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	NewService()
}

func TestWrapValueWei(t *testing.T) {
	type args struct {
		in   *big.Int
		code string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "",
			args: args{
				in:   big.NewInt(1123042123),
				code: "9014",
			},
			want: big.NewInt(1123049014),
		},
		{
			name: "",
			args: args{
				in:   big.NewInt(11),
				code: "9014",
			},
			want: big.NewInt(9014),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrapValueWei(tt.args.in, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapValueWei() = %v, want %v", got, tt.want)
			}
		})
	}
}
