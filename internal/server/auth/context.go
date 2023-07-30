package auth

import (
	"context"
	"time"
)

var ()

type (
	Extractor interface {
		GetUserId(ctx context.Context) (string, error)
	}

	DefaultExtractor struct {
	}

	MockExtractor struct {
	}
	tz struct {
	}
)

func GetTzContext(ctx context.Context) *time.Location {
	v := ctx.Value(tz{})
	if v == nil {
		return time.UTC
	}

	tz, ok := v.(*time.Location)
	if !ok {
		return time.UTC
	}
	return tz
}

func setTzContext(ctx context.Context, location *time.Location) context.Context {
	ctx = context.WithValue(ctx, tz{}, location)
	return ctx
}
