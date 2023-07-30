package repository

import (
	"context"
	"database/sql"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

var ErrNotFound = errors.New("not found")

func (r *pgRepository) GetSettings(ctx context.Context, userId string) (*v1.Settings, error) {
	var payload string
	if err := r.conn.GetContext(ctx, &payload, "select payload from settings where user_id = $1", userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var s v1.Settings
	if err := protojson.Unmarshal([]byte(payload), &s); err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *pgRepository) UpdateSettings(ctx context.Context, request *v1.Settings) error {
	p, err := protojson.Marshal(request)
	if err != nil {
		return err
	}

	s := string(p)

	if _, err := r.conn.ExecContext(ctx, `insert into settings (user_id, payload) values ($1, $2) on conflict (user_id) do update set payload = $2`, request.UserId, &s); err != nil {
		return err
	}
	return nil
}
