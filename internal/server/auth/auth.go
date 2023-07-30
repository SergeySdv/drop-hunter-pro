package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknownUser = errors.New("unknown user")
	ErrNoLocation  = errors.New("location is not set")
)

func AuthGRPC(ctx context.Context) (context.Context, error) {

	token, err := GetToken(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "session not found")
	}

	tz, err := GetTz(ctx)
	if err != nil {
		tz = time.UTC
	}

	u, _, err := ExtractUserAndTokenFromStringToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "session not found")
	}

	ctx = user.SetUserIdContext(ctx, u.UserId)
	ctx = setTzContext(ctx, tz)

	return ctx, nil
}

func GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrUnknownUser
	}

	cookies := md.Get("grpcgateway-cookie")

	token, ok := getCookieWithName(cookies, "session")
	if !ok {
		return "", ErrUnknownUser
	}

	return token, nil
}

func GetTz(ctx context.Context) (*time.Location, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNoLocation
	}

	token := md.Get("tz")
	if len(token) == 0 {
		return nil, ErrNoLocation
	}

	t, err := time.Parse("-07:00", token[0])
	if err != nil {
		return nil, err
	}

	return t.Location(), nil
}

func getCookieWithName(cookies []string, name string) (string, bool) {
	for _, cookie := range cookies {
		pairs := strings.Split(cookie, ";")

		for _, pair := range pairs {

			parts := strings.Split(pair, "=")
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			val := parts[1]

			if key == name {
				return val, true
			}
		}

	}

	return "", false
}
