package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"database/sql"
	"time"

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

	return nil
}

func (a *ProcessArg) ToPB(flowLabel string) (*v1.Process, error) {

	var payload v1.Process
	if err := protojson.Unmarshal([]byte(a.Payload), &payload); err != nil {
		return nil, err
	}

	out := v1.Process{
		Id:         a.Id,
		Status:     v1.ProcessStatus(v1.ProcessStatus_value[a.Status]),
		Profiles:   nil,
		FlowId:     a.FlowId,
		CreatedAt:  timestamppb.New(a.CreatedAt),
		UpdatedAt:  timestamppb.New(a.UpdatedAt),
		FinishedAt: nil,
		StartedAt:  nil,
		FlowLabel:  flowLabel,
	}

	if a.StartedAt.Valid {
		out.StartedAt = timestamppb.New(a.StartedAt.Time)
	}

	if a.FinishedAt.Valid {
		out.FinishedAt = timestamppb.New(a.FinishedAt.Time)
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

func (r *PGRepository) GetProcessUpdatedAt(ctx context.Context, processId string) (*time.Time, error) {
	q := `select updated_at from process where id = $1`
	var a time.Time
	if err := r.conn.GetContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *PGRepository) ListProcessIdsByStatus(ctx context.Context, statuses ...v1.ProcessStatus) ([]string, error) {

	q, args, err := sqlx.In(`select distinct id from process where status in (?)`, statuses)
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

func (r *PGRepository) ListProcessByUser(ctx context.Context, userId string) ([]*v1.Process, error) {

	q := `select * from process where user_id in ($1) order by created_at desc `
	tmp := make([]ProcessArg, 0)
	if err := r.conn.SelectContext(ctx, &tmp, q, userId); err != nil {
		return nil, err
	}

	out := make([]*v1.Process, 0)
	for _, p := range tmp {

		flow, err := r.GetFlow(ctx, &v1.GetFlowRequest{
			Id: p.FlowId,
		})
		if err != nil {
			return nil, err
		}

		pb, err := p.ToPB(flow.Flow.Label)
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

func (r *PGRepository) UpdateProcessStatus(ctx context.Context, req *UpdateProcess) error {

	q := `update process set updated_at = $1, status = $2 where id = $3`
	if _, err := r.conn.ExecContext(ctx, q, time.Now(), req.Status, req.Id); err != nil {
		return err
	}

	return nil
}

func (r *PGRepository) DeleteProcess(ctx context.Context, id string) error {

	process, err := r.GetProcessArg(context.Background(), &v1.GetProcessRequest{Id: id})
	if err != nil {
		return err
	}

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	for _, profile := range process.Process.Profiles {

		for _, task := range profile.Tasks {

			q := `delete from process_task_history where task_id = $1 `
			if _, err := tx.ExecContext(ctx, q, task.Id); err != nil {
				return err
			}

			q = `delete from process_tasks where id = $1 `
			if _, err := tx.ExecContext(ctx, q, task.Id); err != nil {
				return err
			}
		}

		q := `delete from process_profiles where id = $1 `
		if _, err := tx.ExecContext(ctx, q, profile.Id); err != nil {
			return err
		}
	}

	q := `delete from process where id = $1 `
	if _, err := tx.ExecContext(ctx, q, id); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PGRepository) GetProcessStatus(ctx context.Context, processId string) (*v1.ProcessStatus, error) {
	var s string
	if err := r.conn.GetContext(ctx, &s, `select status from process where id = $1`, processId); err != nil {
		return nil, err
	}

	temp := v1.ProcessStatus(v1.ProcessStatus_value[s])
	return &temp, nil
}

func (r *PGRepository) GetProcessUser(ctx context.Context, processId string) (*string, error) {
	var s string
	if err := r.conn.GetContext(ctx, &s, `select user_id from process where id = $1`, processId); err != nil {
		return nil, err
	}

	return &s, nil
}
