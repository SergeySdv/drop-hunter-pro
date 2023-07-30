package repository

import (
	"context"
	"database/sql"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Withdrawer struct {
	Id           string         `db:"id"`
	ExchangeType string         `db:"exchange_type"`
	Label        string         `db:"label"`
	Proxy        sql.NullString `db:"proxy"`
	SecretKey    []byte         `db:"secret_key"`
	ApiKey       []byte         `db:"api_key"`
	UserId       string         `db:"user_id"`
	CreatedAt    time.Time      `db:"created_at"`
	PrevId       sql.NullString `db:"prev_id"`
}

var ExchangeAccountCols = []string{
	"id",
	"exchange_type",
	"label",
	"proxy",
	"secret_key",
	"api_key",
	"user_id",
	"created_at",
	"prev_id",
}

var (
	exacccols = NewHelper(ExchangeAccountCols)
)

func (a *Withdrawer) FromPB(pb *v1.Withdrawer, userId string) {
	a.Id = pb.Id
	a.Label = pb.Label
	a.ExchangeType = pb.ExchangeType.String()

	if pb.Proxy != nil {
		a.Proxy.Valid = true
		a.Proxy.String = *pb.Proxy
	}

	a.CreatedAt = pb.CreatedAt.AsTime()
	a.UserId = userId
}

func (a *Withdrawer) ToPB() *v1.Withdrawer {
	p := &v1.Withdrawer{
		Id:           a.Id,
		ExchangeType: v1.ExchangeType(v1.ExchangeType_value[a.ExchangeType]),
		Label:        a.Label,
		Proxy:        nil,
		CreatedAt:    timestamppb.New(a.CreatedAt),
	}

	if a.Proxy.Valid {
		p.Proxy = &a.Proxy.String
	}

	return p
}

func (r *pgRepository) CreateWithdrawer(ctx context.Context, req *Withdrawer) (*Withdrawer, error) {

	if err := createWithdrawer(ctx, r.conn, req); err != nil {
		return nil, err
	}

	res, err := r.GetWithdrawer(ctx, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *pgRepository) ListSubWithdrawers(ctx context.Context, id, userId string) ([]Withdrawer, error) {
	q := "select * from withdrawers where user_id = $1 and id in (select id from withdrawers where prev_id = $2) order by created_at desc"
	out := make([]Withdrawer, 0)
	if err := r.conn.SelectContext(ctx, &out, q, userId, id); err != nil {
		return nil, err
	}

	return out, nil
}

func (r *pgRepository) CreateSubWithdrawer(ctx context.Context, req *Withdrawer) (err error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
			}
		} else {
			if err := tx.Commit(); err != nil {
			}
		}
	}()

	if err := createWithdrawer(ctx, r.conn, req); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) GetWithdrawer(ctx context.Context, id, userId string) (*Withdrawer, error) {
	return getWithdrawer(ctx, r.conn, id, userId)
}

func getWithdrawer(ctx context.Context, conn *sqlx.DB, id, userId string) (*Withdrawer, error) {
	q := `select * from withdrawers where id = $1 and user_id = $2`
	var a Withdrawer
	if err := conn.GetContext(ctx, &a, q, id, userId); err != nil {
		return nil, err
	}
	return &a, nil
}

func createWithdrawer(ctx context.Context, conn *sqlx.DB, req *Withdrawer) error {
	q := `insert into withdrawers (id, exchange_type, label, proxy, secret_key, api_key, created_at, user_id, prev_id) values 
      (:id, :exchange_type, :label, :proxy, :secret_key, :api_key, :created_at, :user_id, :prev_id)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) ListWithdrawers(ctx context.Context, userId string) ([]Withdrawer, error) {

	res, err := listWithdrawers(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *pgRepository) GetList(ctx context.Context, id string) (*Withdrawer, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	q := "select * from withdrawers where user_id = $1 and id = $2"
	var out Withdrawer
	if err := r.conn.GetContext(ctx, &out, q, userId, id); err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *pgRepository) GetWithdrawers(ctx context.Context, id string) (*Withdrawer, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	q := "select * from withdrawers where user_id = $1 and id = $2"
	var out Withdrawer
	if err := r.conn.GetContext(ctx, &out, q, userId, id); err != nil {
		return nil, err
	}

	return &out, nil
}

func listWithdrawers(ctx context.Context, conn *sqlx.DB, userId string) ([]Withdrawer, error) {
	q := "select * from withdrawers where user_id = $1 and prev_id is null order by created_at desc"
	out := make([]Withdrawer, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	return out, nil
}

type Addr = string

type OkexDepositAddr struct {
	WithdrawerId string `db:"withdrawer_id"`
	Addr         string `db:"okex_deposit_addr"`
	ProfileId    string `db:"profile_id"`
	UserId       string `db:"user_id"`
	ProfileLabel string `db:"profile_label"`
}

func (r *pgRepository) OkexDepositAddrAttach(ctx context.Context, req *OkexDepositAddr) error {
	q := `insert into okex_deposit_addr_profile (withdrawer_id, okex_deposit_addr, profile_id, user_id) 
values (:withdrawer_id, :okex_deposit_addr, :profile_id, :user_id)`
	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) OkexDepositAddrDetach(ctx context.Context, req *OkexDepositAddr) error {
	q := `delete from okex_deposit_addr_profile 
where  withdrawer_id = :withdrawer_id 
    and okex_deposit_addr = :okex_deposit_addr
    and profile_id = :profile_id
     and user_id = :user_id`
	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) ListDepositAddrAttach(ctx context.Context, userId string) (map[Addr]OkexDepositAddr, error) {
	q := `select d.*, cast(p.num as text) as profile_label  from okex_deposit_addr_profile as d 
             join profiles as p on p.id = d.profile_id
             where d.user_id = $1`
	out := make([]OkexDepositAddr, 0)
	if err := r.conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	m := make(map[Addr]OkexDepositAddr)

	for _, o := range out {
		m[o.Addr] = o
	}
	return m, nil
}

func (r *pgRepository) GetAttachedAddr(ctx context.Context, profileId, userId string) (*OkexDepositAddr, error) {
	q := `select * from okex_deposit_addr_profile 
             where user_id = $1 and profile_id = $2`
	var addr OkexDepositAddr
	if err := r.conn.GetContext(ctx, &addr, q, userId, profileId); err != nil {
		return nil, pg.PgError(err)
	}

	return &addr, nil
}

func (r *pgRepository) DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error) {
	if _, err := r.conn.ExecContext(ctx, "delete from withdrawers where id = $1", req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *pgRepository) UpdateWithdrawer(ctx context.Context, req *Withdrawer) error {
	q := `update withdrawers set proxy = :proxy, label = :label where id = :id and user_id = :user_id`

	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) ExportExchangeAccounts(ctx context.Context, userId string) ([]Withdrawer, error) {

	out := make([]Withdrawer, 0)
	q := Join(`select `, exacccols.Cols(), ` from withdrawers where user_id = $1 order by created_at desc `)

	if err := r.conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}
	return out, nil
}
