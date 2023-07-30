package access

import (
	"context"

	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Access(r repository.UserRepository) func(ctx context.Context) (context.Context, error) {

	return func(ctx context.Context) (context.Context, error) {
		id, err := user.GetUserId(ctx)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}
		_, err = r.GetUserById(ctx, id)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "").Err()
		}

		//if !u.Access {
		//	return nil, status.New(codes.PermissionDenied, "").Err()
		//}

		return ctx, nil
	}

}
