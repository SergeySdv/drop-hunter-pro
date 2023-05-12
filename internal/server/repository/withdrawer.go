package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/user"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Withdrawer struct {
	Id           string         `db:"id"`
	ExchangeType string         `db:"exchange_type"`
	Label        string         `db:"label"`
	Proxy        sql.NullString `db:"proxy"`
	SecretKey    string         `db:"secret_key"`
	ApiKey       string         `db:"api_key"`
	UserId       string         `db:"user_id"`
	CreatedAt    time.Time      `db:"created_at"`
}

func (a *Withdrawer) FromPB(pb *v1.Withdrawer, userId string) {
	a.Id = pb.Id
	a.Label = pb.Label
	a.ExchangeType = pb.ExchangeType.String()

	if pb.Proxy != nil {
		a.Proxy.Valid = true
		a.Proxy.String = *pb.Proxy
	}
	a.SecretKey = pb.SecretKey
	a.ApiKey = pb.ApiKey

	a.CreatedAt = pb.CreatedAt.AsTime()
	a.UserId = userId
}

func (a *Withdrawer) ToPB() *v1.Withdrawer {
	p := &v1.Withdrawer{
		Id:           a.Id,
		ExchangeType: v1.ExchangeType(v1.ExchangeType_value[a.ExchangeType]),
		Label:        a.Label,
		SecretKey:    a.SecretKey,
		ApiKey:       a.ApiKey,
		Proxy:        nil,
		CreatedAt:    timestamppb.New(a.CreatedAt),
	}

	if a.Proxy.Valid {
		p.Proxy = &a.Proxy.String
	}

	return p
}

func (r *PGRepository) CreateWithdrawer(ctx context.Context, req *v1.CreateWithdrawerRequest) (*v1.CreateWithdrawerResponse, error) {

	userId := user.GetUserId(ctx)
	pb := &v1.Withdrawer{
		Id:           uuid.New().String(),
		ExchangeType: req.ExchangeType,
		Label:        req.Label,
		SecretKey:    req.SecretKey,
		ApiKey:       req.ApiKey,
		Proxy:        req.Proxy,
		CreatedAt:    timestamppb.Now(),
	}

	a := &Withdrawer{}
	a.FromPB(pb, userId)

	if err := createWithdrawer(ctx, r.conn, a); err != nil {
		return nil, err
	}

	res, err := r.GetWithdrawer(ctx, a.Id)
	if err != nil {
		return nil, err
	}

	return &v1.CreateWithdrawerResponse{
		Withdrawer: res,
	}, nil
}

func (r *PGRepository) GetWithdrawer(ctx context.Context, id string) (*v1.Withdrawer, error) {
	res, err := getWithdrawer(ctx, r.conn, id)
	if err != nil {
		return nil, err
	}
	return res.ToPB(), nil
}

func getWithdrawer(ctx context.Context, conn *sqlx.DB, id string) (*Withdrawer, error) {
	q := `select * from withdrawers where id = $1`
	var a Withdrawer
	if err := conn.GetContext(ctx, &a, q, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func createWithdrawer(ctx context.Context, conn *sqlx.DB, req *Withdrawer) error {
	q := `insert into withdrawers (id, exchange_type, label, proxy, secret_key, api_key, created_at, user_id) values 
      (:id, :exchange_type, :label, :proxy, :secret_key, :api_key, :created_at, :user_id)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *PGRepository) ListWithdrawers(ctx context.Context, req *v1.ListWithdrawerRequest) (*v1.ListWithdrawerResponse, error) {

	userId := user.GetUserId(ctx)

	res, err := listWithdrawers(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return &v1.ListWithdrawerResponse{
		Withdrawers: res,
	}, nil
}

func (r *PGRepository) GetWithdrawers(ctx context.Context, id string) (*Withdrawer, error) {
	userId := user.GetUserId(ctx)

	q := "select * from withdrawers where user_id = $1 and id = $2"
	var out Withdrawer
	if err := r.conn.GetContext(ctx, &out, q, userId, id); err != nil {
		return nil, err
	}

	return &out, nil
}

func listWithdrawers(ctx context.Context, conn *sqlx.DB, userId string) ([]*v1.Withdrawer, error) {
	q := "select * from withdrawers where user_id = $1 order by created_at desc"
	out := make([]Withdrawer, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	tmp := make([]*v1.Withdrawer, 0)
	for i := range out {
		tmp = append(tmp, out[i].ToPB())
	}

	return tmp, nil
}
func (r *PGRepository) DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error) {
	if _, err := r.conn.ExecContext(ctx, "delete from withdrawers where id = $1", req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}
