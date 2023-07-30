package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProcessArg struct {
	Process
	Profiles []ProcessProfileArg
}

type Process struct {
	Id         string       `db:"id"`
	Status     string       `db:"status"`
	UserId     string       `db:"user_id"`
	FlowId     string       `db:"flow_id"`
	Payload    string       `db:"payload"`
	UpdatedAt  time.Time    `db:"updated_at"`
	CreatedAt  time.Time    `db:"created_at"`
	StartedAt  sql.NullTime `db:"started_at"`
	FinishedAt sql.NullTime `db:"finished_at"`
	Progress   int          `db:"progress"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	AutoRetry  bool         `db:"auto_retry"`
}

func (a *ProcessArg) FromPB(pb *v1.Process, userId string) error {

	// process
	a.Id = pb.Id
	a.Status = pb.Status.String()
	a.UserId = userId
	a.FlowId = pb.FlowId
	a.UpdatedAt = pb.UpdatedAt.AsTime()

	profiles := make([]ProcessProfileArg, 0)
	for _, p := range pb.Profiles {
		var profile ProcessProfileArg
		if err := profile.FromPB(p, a.Id); err != nil {
			return err
		}
		profiles = append(profiles, profile)
	}

	a.Profiles = profiles

	s, err := protojson.Marshal(pb)
	if err != nil {
		return err
	}
	a.Payload = string(s)

	if pb.StartedAt != nil {
		a.StartedAt.Valid = true
		a.StartedAt.Time = pb.StartedAt.AsTime()
	}

	if pb.FinishedAt != nil {
		a.FinishedAt.Valid = true
		a.FinishedAt.Time = pb.FinishedAt.AsTime()
	}

	if pb.DeletedAt != nil {
		a.DeletedAt.Valid = true
		a.DeletedAt.Time = pb.DeletedAt.AsTime()
	}

	a.AutoRetry = pb.AutoRetry

	return nil
}

func (a *ProcessArg) ToPB(flow *v1.Flow) (*v1.Process, error) {

	var payload v1.Process
	if err := protojson.Unmarshal([]byte(a.Payload), &payload); err != nil {
		return nil, err
	}

	out := v1.Process{
		Id:         a.Id,
		Status:     v1.ProcessStatus(v1.ProcessStatus_value[a.Status]),
		FlowId:     a.FlowId,
		CreatedAt:  timestamppb.New(a.CreatedAt),
		UpdatedAt:  timestamppb.New(a.UpdatedAt),
		FinishedAt: nil,
		StartedAt:  nil,
		FlowLabel:  flow.Label,
		Progress:   int64(a.Progress),
		DeletedAt:  nil,
		AutoRetry:  a.AutoRetry,
		Flow:       flow,
	}

	if a.StartedAt.Valid {
		out.StartedAt = timestamppb.New(a.StartedAt.Time)
	}

	if a.FinishedAt.Valid {
		out.FinishedAt = timestamppb.New(a.FinishedAt.Time)
	}

	if a.DeletedAt.Valid {
		out.DeletedAt = timestamppb.New(a.DeletedAt.Time)
	}

	pr := make([]*v1.ProcessProfile, 0)
	for _, p := range a.Profiles {
		profilePB, err := p.ToPB()
		if err != nil {
			return nil, err
		}

		pr = append(pr, profilePB)
	}
	out.Profiles = pr

	return &out, nil
}

func (r *pgRepository) GetProcessProgress(ctx context.Context, processId string) (*int, error) {
	q := `select round(100 *
((select count(*) from process_profiles as pp
                left join process_tasks pt on pp.id = pt.profile_id
                where pp.process_id = $1
and pt.status = 'StatusDone')::numeric
    /

(select count(*) from process_profiles as pp
                left join process_tasks pt on pp.id = pt.profile_id
                where pp.process_id = $1))::numeric);`
	var a int
	if err := r.conn.GetContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) GetProcessUpdatedAt(ctx context.Context, processId string) (*time.Time, error) {
	q := `select updated_at from process where id = $1`
	var a time.Time
	if err := r.conn.GetContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) ListProcessIdsByStatus(ctx context.Context, statuses ...v1.ProcessStatus) ([]string, error) {

	q, args, err := sqlx.In(`select distinct id from process where status in (?) and deleted_at is null`, statuses)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)
	tmp := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &tmp, q, args...); err != nil {
		return nil, err
	}

	return tmp, nil
}

func (r *pgRepository) ListProcessIdsForAutoRetry(ctx context.Context) ([]string, error) {

	q, args, err := sqlx.In(`select distinct id from process where status in (?) and deleted_at is null and auto_retry = true`, v1.ProcessStatus_StatusError.String())
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)
	tmp := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &tmp, q, args...); err != nil {
		return nil, err
	}

	return tmp, nil
}

func (r *pgRepository) ListProcessByUser(ctx context.Context, userId string, statuses []string, offset, limit int) ([]*v1.Process, error) {

	q := `select * from process where 
                          user_id = $1
                          and deleted_at is null `

	arg := []interface{}{userId}
	count := 2
	if len(statuses) > 0 {
		base := "("

		for _, s := range statuses {
			base += fmt.Sprintf(`$%d, `, count)
			arg = append(arg, s)
			count++
		}
		base = base[:len(base)-2]

		q = q + " and status in " + base + ") "
	}

	q = q + fmt.Sprintf("order by created_at desc limit $%d ", count)
	count++
	q = q + fmt.Sprintf("offset $%d ", count)

	arg = append(arg, limit, offset)

	tmp := make([]ProcessArg, 0)
	if err := r.conn.SelectContext(ctx, &tmp, q, arg...); err != nil {
		return nil, err
	}

	out := make([]*v1.Process, 0)
	for i := range tmp {
		p := &tmp[i]

		flow, err := r.GetFlow(ctx, &v1.GetFlowRequest{
			Id: p.FlowId,
		})
		if err != nil {
			return nil, err
		}

		progress, err := r.GetProcessProgress(ctx, p.Id)
		if err != nil {
			return nil, err
		}
		p.Progress = *progress

		pps, err := r.GetProcessProfiles(ctx, p.Id)
		if err != nil {
			return nil, err
		}

		profiles := make([]ProcessProfileArg, 0)

		for _, pp := range pps {
			profiles = append(profiles, ProcessProfileArg{
				ProcessProfile: pp,
				Tasks:          nil,
			})
		}

		p.Profiles = profiles

		pb, err := p.ToPB(flow.Flow)
		if err != nil {
			return nil, err
		}
		out = append(out, pb)
	}

	return out, nil
}

type Valid[T any] struct {
	Valid bool
	T     any
}

type UpdateProcess struct {
	Id     string
	Status string
}

func (r *pgRepository) UpdateProcessAutoRetry(ctx context.Context, id string, enable bool) error {
	q := `update process set auto_retry = $1 where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, enable, id); err != nil {
		return err
	}

	if err := r.UpdateProcessTime(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) UpdateProcessTime(ctx context.Context, id string) error {
	q := `update process set updated_at = $1 where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, time.Now(), id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) UpdateProcessStatus(ctx context.Context, req *UpdateProcess) error {

	q := `update process set updated_at = $1, status = $2 where id = $3`
	if _, err := r.conn.ExecContext(ctx, q, time.Now(), req.Status, req.Id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) DeleteProcess(ctx context.Context, id string) error {
	q := `update process set deleted_at = now() where id = $1 `
	if _, err := r.conn.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) GetProcessStatus(ctx context.Context, processId string) (*v1.ProcessStatus, error) {
	var s string
	if err := r.conn.GetContext(ctx, &s, `select status from process where id = $1`, processId); err != nil {
		return nil, err
	}

	temp := v1.ProcessStatus(v1.ProcessStatus_value[s])
	return &temp, nil
}

func (r *pgRepository) GetProcessUser(ctx context.Context, processId string) (*string, error) {
	var s string
	if err := r.conn.GetContext(ctx, &s, `select user_id from process where id = $1`, processId); err != nil {
		return nil, err
	}

	return &s, nil
}
