package repository

import (
	"context"
	"crypto_scripts/internal/defi"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/user"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Profile struct {
	Id              string         `db:"id"`
	Label           string         `db:"label"`
	Proxy           sql.NullString `db:"proxy"`
	MmskPk          string         `db:"mmsk_pk"`
	Meta            sql.NullString `db:"meta"`
	UserId          string         `db:"user_id"`
	CreatedAt       time.Time      `db:"created_at"`
	OkexDepositAddr sql.NullString `db:"okex_deposit_addr"`
	OkexAccName     sql.NullString `db:"okex_acc_name"`
}

func (a *Profile) ToPB() *v1.Profile {

	transactor, err := defi.NewWalletTransactor(a.MmskPk)
	if err != nil {
		return nil
	}

	p := &v1.Profile{
		Id:        a.Id,
		Label:     a.Label,
		Proxy:     nil,
		MmskId:    transactor.WalletAddrHR,
		Meta:      nil,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	if a.Proxy.Valid {
		p.Proxy = &a.Proxy.String
	}

	if a.Meta.Valid {
		p.Meta = &a.Meta.String
	}

	if a.OkexDepositAddr.Valid || a.OkexAccName.Valid {
		p.OkexAccount = &v1.OkexAccount{
			SubAccName:  a.OkexAccName.String,
			DepositAddr: a.OkexDepositAddr.String,
		}
	}

	return p
}

func (r *PGRepository) CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error) {

	userId := user.GetUserId(ctx)

	a := &Profile{
		Id:        uuid.New().String(),
		Label:     req.Label,
		Proxy:     sql.NullString{},
		MmskPk:    req.MmskPk,
		Meta:      sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	if req.Meta != nil {
		a.Meta.Valid = true
		a.Meta.String = *req.Meta
	}

	if req.Proxy != nil {
		a.Proxy.Valid = true
		a.Proxy.String = *req.Proxy
	}

	if req.OkexAccount != nil {
		if req.OkexAccount.SubAccName != "" {
			a.OkexAccName.Valid = true
			a.OkexAccName.String = req.OkexAccount.SubAccName
		}
		if req.OkexAccount.DepositAddr != "" {
			a.OkexDepositAddr.Valid = true
			a.OkexDepositAddr.String = req.OkexAccount.DepositAddr
		}
	}

	if err := CreateProfile(ctx, r.conn, a); err != nil {
		return nil, err
	}

	res, err := r.GetProfile(ctx, a.Id)
	if err != nil {
		return nil, err
	}

	return &v1.CreateProfileResponse{
		Profile: res.ToPB(),
	}, nil
}

func CreateProfile(ctx context.Context, conn *sqlx.DB, req *Profile) error {
	q := `insert into profiles (id, label, proxy, mmsk_pk, meta, created_at, user_id, okex_acc_name, okex_acc_name) values 
      (:id, :label, :proxy, :mmsk_pk, :meta, :created_at, :user_id, :okex_acc_name, :okex_acc_name)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *PGRepository) GetProfile(ctx context.Context, id string) (*Profile, error) {
	return GetProfile(ctx, r.conn, id)
}

func (r *PGRepository) UpdateProfile(ctx context.Context, req *Profile) error {
	q := `update profiles set
			okex_deposit_addr = :okex_deposit_addr,
			okex_acc_name = :okex_acc_name,
			meta = :meta,
			proxy = :proxy,
			label = :label
		where id = :id
                    `

	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *PGRepository) ValidateLabel(ctx context.Context, request *v1.ValidateLabelRequest) (*bool, error) {
	if request.ProfileId == nil {
		q := `select count(*) from profiles where label = $1`
		var a int
		if err := r.conn.GetContext(ctx, &a, q, request.Label); err != nil {
			return nil, err
		}
		b := a == 0
		return &b, nil
	} else {
		q := `select count(*) from profiles where label = $1 and id not in ($2)`
		var a int
		if err := r.conn.GetContext(ctx, &a, q, request.Label, request.GetProfileId()); err != nil {
			return nil, err
		}
		b := a == 0
		return &b, nil
	}
}

func GetProfile(ctx context.Context, conn *sqlx.DB, id string) (*Profile, error) {
	q := `select * from profiles where id = $1`
	var a Profile
	if err := conn.GetContext(ctx, &a, q, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *PGRepository) ListProfiles(ctx context.Context, req *v1.ListProfileRequest) (*v1.ListProfileResponse, error) {

	userId := user.GetUserId(ctx)

	res, err := listProfiles(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return &v1.ListProfileResponse{
		Profiles: res,
	}, nil
}

func listProfiles(ctx context.Context, conn *sqlx.DB, userId string) ([]*v1.Profile, error) {
	q := "select * from profiles where user_id = $1 order by created_at desc"
	out := make([]Profile, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	tmp := make([]*v1.Profile, 0)
	for i := range out {
		tmp = append(tmp, out[i].ToPB())
	}

	return tmp, nil
}

func (r *PGRepository) SearchProfile(ctx context.Context, req *v1.SearchProfileRequest) (*v1.SearchProfileResponse, error) {
	q := `select * from profiles where label like '%` + req.Pattern + `%' limit 10`
	out := make([]Profile, 0)
	if err := r.conn.SelectContext(ctx, &out, q); err != nil {
		return nil, err
	}

	tmp := make([]*v1.Profile, 0)
	for i := range out {
		tmp = append(tmp, out[i].ToPB())
	}

	return &v1.SearchProfileResponse{
		Profiles: tmp,
	}, nil
}

func (r *PGRepository) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {
	if _, err := r.conn.ExecContext(ctx, "delete from profiles where id = $1", req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}
