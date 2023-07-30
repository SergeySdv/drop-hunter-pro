package repository

import (
	"context"
)

func (r *pgRepository) GetDailyUserImpact(ctx context.Context, userId string) (*int64, error) {
	q := `select count(1) from transactions where user_id = $1 and created_at > (now() - INTERVAL '1 days');`
	var i int64
	if err := r.conn.GetContext(ctx, &i, q, userId); err != nil {
		return nil, err
	}
	return &i, nil
}
func (r *pgRepository) GetDailyTopImpact(ctx context.Context) (*int64, error) {
	q := `select coalesce(max(uc), 0) from (
select count(1) as uc from transactions where created_at > (now() - INTERVAL '1 days') group by user_id
                                                                                              ) as bb`
	var i int64
	if err := r.conn.GetContext(ctx, &i, q); err != nil {
		return nil, err
	}
	return &i, nil
}
func (r *pgRepository) GetDailyTotalImpact(ctx context.Context) (*int64, error) {
	q := `select count(1) from transactions where created_at > (now() - INTERVAL '1 days');`
	var i int64
	if err := r.conn.GetContext(ctx, &i, q); err != nil {
		return nil, err
	}
	return &i, nil
}
