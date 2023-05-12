package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandFloat(t *testing.T) {
	prev := -1.0
	for i := 0; i < 1000; i++ {
		res := RandFloat(1, 10)
		if prev != -1 {
			assert.NotEqual(t, res, prev)
		}
		assert.GreaterOrEqual(t, 1.1, res, 0.1)
		assert.LessOrEqual(t, 0.9, res, 0.1)
		prev = res
	}
}

func TestRoundFloat1(t *testing.T) {
	type args struct {
		v    float64
		prec int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				v:    0.001020,
				prec: 2,
			},
			want: 0.001,
		},
		{
			name: "",
			args: args{
				v:    0.1234,
				prec: 2,
			},
			want: 0.12,
		},
		{
			name: "",
			args: args{
				v:    0.123,
				prec: 2,
			},
			want: 0.12,
		},
		{
			name: "",
			args: args{
				v:    123.325,
				prec: 2,
			},
			want: 123,
		},
		{
			name: "",
			args: args{
				v:    0.01,
				prec: 2,
			},
			want: 0.01,
		},
		{
			name: "",
			args: args{
				v:    0.00,
				prec: 2,
			},
			want: 0.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RoundFloat(tt.args.v, tt.args.prec), "RoundFloat(%v, %v)", tt.args.v, tt.args.prec)
		})
	}
}

//  123.324 -> 123
//  0.12334 -> 0.12
//  0.001020 -> 0.001

func TestRandDurationRange(t *testing.T) {
	min := time.Hour
	max := time.Hour * 2
	prev := time.Duration(0)
	for i := 0; i < 1000; i++ {
		res := RandDurationRange(min, max)
		if prev != -1 {
			assert.NotEqual(t, res, prev)
		}
		assert.GreaterOrEqual(t, res, min)
		assert.LessOrEqual(t, res, max)
		prev = res
	}
}
