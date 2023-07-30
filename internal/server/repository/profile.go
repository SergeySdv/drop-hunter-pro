package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Profile struct {
	Id        string         `db:"id"`
	Label     string         `db:"label"`
	Proxy     sql.NullString `db:"proxy"`
	MmskPk    []byte         `db:"mmsk_pk"`
	Meta      sql.NullString `db:"meta"`
	UserId    string         `db:"user_id"`
	CreatedAt time.Time      `db:"created_at"`
	Num       int64          `db:"num"`
	UserAgent string         `db:"user_agent"`
	MmskId    []byte         `db:"mmsk_id"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

var ProfileCols = []string{
	"id",
	"label",
	"proxy",
	"mmsk_pk",
	"meta",
	"user_id",
	"created_at",
	"num",
	"user_agent",
	"mmsk_id",
	"deleted_at",
}

var (
	prfh = NewHelper(ProfileCols)
)

func (a *Profile) ToPB() (*v1.Profile, error) {

	transactor, err := defi.NewWalletTransactor(string(a.MmskPk))
	if err != nil {
		return nil, err
	}

	p := &v1.Profile{
		Id:        a.Id,
		Label:     a.Label,
		Proxy:     nil,
		MmskId:    transactor.WalletAddrHR,
		Meta:      nil,
		CreatedAt: timestamppb.New(a.CreatedAt),
		UserAgent: a.UserAgent,
		Num:       a.Num,
	}

	if a.Proxy.Valid {
		p.Proxy = &a.Proxy.String
	}

	if a.Meta.Valid {
		p.Meta = &a.Meta.String
	}

	if a.DeletedAt.Valid {
		p.DeletedAt = timestamppb.New(a.DeletedAt.Time)
	}

	return p, nil
}

func (r *pgRepository) CreateProfile(ctx context.Context, req *Profile) (err error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	txx := pg.WrapSqlxTx(tx)

	maxNum, err := SelectProfileMaxNumByUser(ctx, txx, req.UserId)
	if err != nil {
		return err
	}
	maxNum++
	req.Num = maxNum

	if err := CreateProfile(ctx, txx, req); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func SelectProfileMaxNumByUser(ctx context.Context, conn pg.SqlDriver, userId string) (int64, error) {
	q := `select coalesce(max(num), 0) from profiles where user_id = $1`
	var num int64 = 1
	if err := conn.GetContext(ctx, &num, q, userId); err != nil {
		return 0, err
	}
	return num, nil
}
func CreateProfile(ctx context.Context, conn pg.SqlDriver, req *Profile) error {
	q := Join("insert into profiles (", prfh.Cols(), `) values (`, prfh.ColsColon(), ")")

	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return pg.PgError(err)
	}
	return nil
}
func (r *pgRepository) GetProfile(ctx context.Context, id string) (*Profile, error) {
	return GetProfile(ctx, r.conn, id)
}
func (r *pgRepository) UpdateProfile(ctx context.Context, req *Profile) error {
	q := `update profiles set
			meta = :meta,
			proxy = :proxy,
			label = :label,
			user_agent = :user_agent
		where id = :id`

	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

type ValidateLabelReq struct {
	ProfileId *string
	Label     string
	UserId    string
}

func (r *pgRepository) ValidateLabel(ctx context.Context, request *ValidateLabelReq) (*bool, error) {
	if request.ProfileId == nil {
		q := `select count(*) from profiles where label = $1 and user_id = $2 and deleted_at is null`
		var a int
		if err := r.conn.GetContext(ctx, &a, q, request.Label, request.UserId); err != nil {
			return nil, err
		}
		b := a == 0
		return &b, nil
	} else {
		q := `select count(*) from profiles where label = $1 and id not in ($2) and user_id = $3 and deleted_at is null`
		var a int
		if err := r.conn.GetContext(ctx, &a, q, request.Label, *request.ProfileId, request.UserId); err != nil {
			return nil, err
		}
		b := a == 0
		return &b, nil
	}
}
func GetProfile(ctx context.Context, conn *sqlx.DB, id string) (*Profile, error) {
	q := Join(`select `, prfh.Cols(), ` from profiles where id = $1`)
	var a Profile
	if err := conn.GetContext(ctx, &a, q, id); err != nil {
		return nil, err
	}
	return &a, nil
}
func (r *pgRepository) ListProfiles(ctx context.Context, userId string) ([]Profile, error) {

	res, err := listProfiles(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func listProfiles(ctx context.Context, conn *sqlx.DB, userId string) ([]Profile, error) {
	q := Join(`select `, prfh.Cols(), ` from profiles where user_id = $1 and deleted_at is null order by num asc`)
	out := make([]Profile, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	return out, nil
}
func (r *pgRepository) SearchNotConnectedOkexDepositProfile(ctx context.Context, userId string) ([]Profile, error) {
	q := Join(`select `, prfh.ColsPref(), ` from profiles as p 
where p.id not in (select profile_id from okex_deposit_addr_profile) and p.user_id = $1 and p.deleted_at is null order by num asc limit 100`)
	out := make([]Profile, 0)
	if err := r.conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	return out, nil
}
func (r *pgRepository) SearchProfile(ctx context.Context, pattern, userId string) ([]Profile, error) {

	q := ``
	out := make([]Profile, 0)

	if pattern == "" || pattern == "*" {
		q := Join(`select `, prfh.Cols(), ` from profiles where user_id = $1 and deleted_at is null order by num asc limit 10`)

		if err := r.conn.SelectContext(ctx, &out, q, userId); err != nil {
			return nil, err
		}
	} else if strings.Contains(pattern, "-") {
		sub := strings.Split(pattern, "-")
		if len(sub) < 2 {
			return nil, errors.New("mask is D-D")
		}

		from, err := strconv.Atoi(sub[0])
		if err != nil {
			return nil, errors.New("mask is D-D")
		}

		to, err := strconv.Atoi(sub[1])
		if err != nil {
			return nil, errors.New("mask is D-D")
		}

		q := Join(`select `, prfh.Cols(), ` from profiles where user_id = $1 and num >= $2 and num <= $3 and deleted_at is null order by num asc limit 100`)

		if err := r.conn.SelectContext(ctx, &out, q, userId, from, to); err != nil {
			return nil, err
		}
	} else {
		num, err := strconv.Atoi(pattern)
		if err != nil {
			return nil, errors.New("mask is D")
		}

		q = Join(`select `, prfh.Cols(), ` from profiles where user_id = $1 and num = $2 and deleted_at is null`)

		if err := r.conn.SelectContext(ctx, &out, q, userId, num); err != nil {
			return nil, err
		}
	}

	return out, nil
}
func (r *pgRepository) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {
	if _, err := r.conn.ExecContext(ctx, "update profiles set deleted_at = now() where id = $1", req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *pgRepository) ExportProfiles(ctx context.Context, userId string) ([]Profile, error) {

	out := make([]Profile, 0)
	q := Join(`select `, prfh.Cols(), ` from profiles where user_id = $1 order by num asc `)

	if err := r.conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}
	return out, nil
}
