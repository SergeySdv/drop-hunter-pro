package user

import (
	"context"

	"github.com/pkg/errors"
)

type userIdKey struct {
}

var ErrUserNotFound = errors.New("user not found in context")

func SetUserIdContext(ctx context.Context, id string) context.Context {
	ctx = context.WithValue(ctx, &userIdKey{}, id)
	return ctx
}

func GetUserId(ctx context.Context) (string, error) {
	val := ctx.Value(&userIdKey{})
	if val == nil {
		return "", ErrUserNotFound
	}

	u, ok := val.(string)
	if !ok {
		return "", ErrUserNotFound
	}

	return u, nil
}
