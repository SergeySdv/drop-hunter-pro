package lib

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testReq struct {
	retry int
}

type testRes struct {
}

var (
	er = errors.New("gg")
	i  = 0
)

func TestRetry(t *testing.T) {
	ctx := context.Background()

	t.Run("fail", func(t *testing.T) {
		i = 0
		ret := 10
		res, err := Retry[*testReq, *testRes](ctx, testFunc, &testReq{ret}, &RetryOpt{
			RetryCount: 2,
		})
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("success", func(t *testing.T) {
		i = 0
		ret := 3
		res, err := Retry[*testReq, *testRes](ctx, testFunc, &testReq{ret - 1}, &RetryOpt{
			RetryCount: ret,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func testFunc(ctx context.Context, arg *testReq) (*testRes, error) {
	if i >= arg.retry {
		return &testRes{}, nil
	}
	i++
	return nil, er
}
