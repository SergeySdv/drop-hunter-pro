package repository

import (
	"context"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProcessTask struct {
	Id        string `db:"id"`
	ProcessId string `db:"process_id"`
	ProfileId string `db:"profile_id"`
	Status    string `db:"status"`
	Payload   string `db:"payload"`
	Weight    int64  `db:"weight"`
}

func (a *ProcessTask) FromPB(t *v1.ProcessTask, processId, profileId string) error {
	a.Status = t.Status.String()
	a.ProcessId = processId
	a.ProfileId = profileId
	a.Id = t.Id
	s, err := protojson.Marshal(t)
	if err != nil {
		return err
	}
	a.Weight = t.Task.Weight
	a.Payload = string(s)

	return nil
}

func (a *ProcessTask) ToPB() (*v1.ProcessTask, error) {

	var payload v1.ProcessTask
	if err := protojson.Unmarshal([]byte(a.Payload), &payload); err != nil {
		return nil, err
	}

	out := v1.ProcessTask{
		Id:           a.Id,
		Task:         payload.Task,
		Status:       v1.ProcessStatus(v1.ProcessStatus_value[a.Status]),
		Transactions: payload.Transactions,
		FinishedAt:   payload.FinishedAt,
		StartedAt:    payload.StartedAt,
		Error:        payload.Error,
	}

	return &out, nil
}

func (r *pgRepository) GetTaskStatus(ctx context.Context, taskId string) (*v1.ProcessStatus, error) {
	var s string
	if err := r.conn.GetContext(ctx, &s, `select status from process_tasks where id = $1`, taskId); err != nil {
		return nil, err
	}

	temp := v1.ProcessStatus(v1.ProcessStatus_value[s])
	return &temp, nil
}
func (r *pgRepository) UpdateProcessTaskStatus(ctx context.Context, status, taskId, processId string) error {

	updatedAt := time.Now()

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	q := `update process_tasks set status = $1  where id = $2`
	if _, err := tx.ExecContext(ctx, q, status, taskId); err != nil {
		return err
	}

	q = `update process set updated_at = $1 where id = $2`
	if _, err := tx.ExecContext(ctx, q, updatedAt, processId); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) UpdateProcessTask(ctx context.Context, req *v1.ProcessTask, taskId, processId, profileId string) error {

	updatedAt := time.Now()
	var p ProcessTask
	if err := p.FromPB(req, processId, profileId); err != nil {
		return err
	}

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	q := `update process_tasks set payload = $1, status = $2  where id = $3`
	if _, err := r.conn.ExecContext(ctx, q, p.Payload, p.Status, taskId); err != nil {
		return err
	}

	q = `update process set updated_at = $1 where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, updatedAt, processId); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) GetProcessProfileTaskStatuses(ctx context.Context, profileId string) ([]v1.ProcessStatus, error) {
	q := `select status from process_tasks where profile_id = $1`
	temp := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &temp, q, profileId); err != nil {
		return nil, err
	}

	statuses := make([]v1.ProcessStatus, 0)

	for _, statusString := range temp {
		statuses = append(statuses, v1.ProcessStatus(v1.ProcessStatus_value[statusString]))
	}

	return statuses, nil
}
