package lib

import (
	"context"
	"strconv"
	"time"

	"github.com/hardstylez72/cry/internal/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RetryOpt struct {
	RetryCount int
	Delay      time.Duration
}

func Retry[input any, output any](ctx context.Context, f func(context.Context, input) (output, error), arg input, opt *RetryOpt) (ret output, err error) {

	if opt == nil {
		opt = &RetryOpt{}
	}

	retries := opt.RetryCount
	for retries >= 0 {
		ret, err = f(ctx, arg)
		if err == nil {
			return ret, nil
		}

		log.Log.Warn("Retry. attempt: "+strconv.Itoa(opt.RetryCount-retries), zap.Error(err))
		retries--

		time.Sleep(opt.Delay)
	}

	return ret, errors.Wrap(err, "retry limit exceeded: "+strconv.Itoa(opt.RetryCount)+" attempts")
}
