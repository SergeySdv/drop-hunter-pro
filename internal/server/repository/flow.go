package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flow struct {
	Id        string         `db:"id"`
	Label     string         `db:"label"`
	Payload   string         `db:"payload"`
	NextId    sql.NullString `db:"next_id"`
	UserId    string         `db:"user_id"`
	CreatedAt time.Time      `db:"created_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

func (a *Flow) FromPB(pb *v1.Flow, userId string) error {
	a.Id = pb.Id
	a.Label = pb.Label
	b, err := protojson.Marshal(pb)
	if err != nil {
		return err
	}
	a.Payload = string(b)
	a.CreatedAt = pb.CreatedAt.AsTime()
	a.UserId = userId
	if pb.DeletedAt != nil {
		a.DeletedAt.Valid = true
		a.DeletedAt.Time = pb.DeletedAt.AsTime()
	}

	return nil
}

func (a *Flow) ToPB() (*v1.Flow, error) {

	var f v1.Flow
	if err := protojson.Unmarshal([]byte(a.Payload), &f); err != nil {
		return nil, err
	}
	p := &v1.Flow{
		Id:        a.Id,
		Label:     a.Label,
		Tasks:     f.Tasks,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	if a.NextId.Valid {
		p.NextId = &a.NextId.String
	}

	if a.DeletedAt.Valid {
		p.DeletedAt = timestamppb.New(a.DeletedAt.Time)
	}

	return p, nil
}

func (r *pgRepository) UpdateFlow(ctx context.Context, parentFlowId string, req *Flow) (err error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}

	}()

	txx := pg.WrapSqlxTx(tx)

	if err := createFlow(ctx, r.conn, req); err != nil {
		return err
	}
	if err := updateFlowParent(ctx, txx, parentFlowId, req.Id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) CreateFlow(ctx context.Context, req *v1.CreateFlowRequest) (*v1.CreateFlowResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	pb := &v1.Flow{
		Id:        uuid.New().String(),
		Label:     req.Label,
		Tasks:     req.Tasks,
		CreatedAt: timestamppb.Now(),
	}

	a := &Flow{}
	if err := a.FromPB(pb, userId); err != nil {
		return nil, err
	}

	if err := createFlow(ctx, r.conn, a); err != nil {
		return nil, err
	}

	res, err := r.GetFlow(ctx, &v1.GetFlowRequest{Id: a.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CreateFlowResponse{
		Flow: res.Flow,
	}, nil
}

func updateFlowParent(ctx context.Context, driver pg.SqlDriver, flowId, nextFlowId string) error {
	q := `update flow set next_id = $1 where id = $2                                                                 `
	if _, err := driver.ExecContext(ctx, q, nextFlowId, flowId); err != nil {
		return err
	}
	return nil
}

func createFlow(ctx context.Context, driver pg.SqlDriver, req *Flow) error {
	q := `insert into flow (id, label, payload, created_at, user_id, next_id) values 
      (:id, :label, :payload, :created_at, :user_id, null)                                                                 `
	if _, err := driver.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) GetFlow(ctx context.Context, req *v1.GetFlowRequest) (*v1.GetFlowResponse, error) {
	res, err := getFlow(ctx, r.conn, req.Id)
	if err != nil {
		return nil, err
	}

	t, err := res.ToPB()
	if err != nil {
		return nil, err
	}
	return &v1.GetFlowResponse{
		Flow: t,
	}, nil
}

func getFlow(ctx context.Context, conn *sqlx.DB, id string) (*Flow, error) {
	q := `select * from flow where id = $1`
	var a Flow
	if err := conn.GetContext(ctx, &a, q, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) ListFlows(ctx context.Context, req *v1.ListFlowRequest) (*v1.ListFlowResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := listFlows(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	return &v1.ListFlowResponse{
		Flows: res,
	}, nil
}

func listFlows(ctx context.Context, conn *sqlx.DB, userId string) ([]*v1.Flow, error) {
	q := "select * from flow where user_id = $1 and next_id is null and deleted_at is null order by created_at desc"
	out := make([]Flow, 0)
	if err := conn.SelectContext(ctx, &out, q, userId); err != nil {
		return nil, err
	}

	tmp := make([]*v1.Flow, 0)
	for i := range out {

		p, err := out[i].ToPB()
		if err != nil {
			return nil, err
		}

		tmp = append(tmp, p)
	}

	return tmp, nil
}
func (r *pgRepository) DeleteFlow(ctx context.Context, req *v1.DeleteFlowRequest) (*v1.DeleteFlowResponse, error) {
	if _, err := r.conn.ExecContext(ctx, "update flow set deleted_at = now() where id = $1", req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}
