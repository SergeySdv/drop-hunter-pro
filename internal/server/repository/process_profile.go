package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/user"
	"time"
)

type ProcessProfileArg struct {
	ProcessProfile
	Tasks []ProcessTask
}

type ProcessProfile struct {
	Id        string `db:"id"`
	ProcessId string `db:"process_id"`
	ProfileId string `db:"profile_id"`
	Status    string `db:"status"`
	Weight    int64  `db:"weight"`
}

func (a *ProcessProfileArg) FromPB(t *v1.ProcessProfile, processId string) error {
	a.Status = t.Status.String()
	a.ProcessId = processId
	a.ProfileId = t.ProfileId

	a.Id = t.Id
	a.Weight = t.Weight
	// tasks
	tasks := make([]ProcessTask, 0)
	for _, t := range t.Tasks {
		var task ProcessTask
		if err := task.FromPB(t, processId, a.Id); err != nil {
			return err
		}
		tasks = append(tasks, task)
	}

	a.Tasks = tasks

	return nil
}

func (a *ProcessProfileArg) ToPB() (*v1.ProcessProfile, error) {

	out := &v1.ProcessProfile{
		Id:        a.Id,
		ProfileId: a.ProfileId,
		Weight:    a.Weight,
		Tasks:     nil,
		Status:    v1.ProcessStatus(v1.ProcessStatus_value[a.Status]),
	}

	tt := make([]*v1.ProcessTask, 0)
	for _, t := range a.Tasks {
		pb, err := t.ToPB()
		if err != nil {
			return nil, err
		}

		tt = append(tt, pb)
	}

	out.Tasks = tt

	return out, nil
}

func (r *PGRepository) UpdateProcessProfile(ctx context.Context, req *v1.ProcessProfile, processId string) (*v1.ProcessProfile, error) {

	updatedAt := time.Now()
	var p ProcessProfileArg
	if err := p.FromPB(req, user.GetUserId(ctx)); err != nil {
		return nil, err
	}

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	q := `update process_profiles set status = $1  where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, p.Status, req.Id); err != nil {
		return nil, err
	}

	q = `update process set updated_at = $1 where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, updatedAt, processId); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return p.ToPB()
}

func (r *PGRepository) GetProcessProfileIds(ctx context.Context, processId string) ([]string, error) {
	q := `select id from process_profiles where process_id = $1 order by weight asc`
	a := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *PGRepository) GetProcessProfileStatuses(ctx context.Context, processId string) ([]v1.ProcessStatus, error) {
	q := `select status from process_profiles where process_id = $1`
	temp := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &temp, q, processId); err != nil {
		return nil, err
	}

	statuses := make([]v1.ProcessStatus, 0)

	for _, statusString := range temp {
		statuses = append(statuses, v1.ProcessStatus(v1.ProcessStatus_value[statusString]))
	}

	return statuses, nil
}

func (r *PGRepository) GetProcessProfileArg(ctx context.Context, profileId string) (*ProcessProfileArg, error) {

	profile, err := r.GetProcessProfile(ctx, profileId)
	if err != nil {
		return nil, err
	}
	tasks, err := r.GetProcessTasks(ctx, profileId)
	if err != nil {
		return nil, err
	}

	p := &ProcessProfileArg{
		ProcessProfile: *profile,
		Tasks:          tasks,
	}
	return p, nil
}

func (r *PGRepository) GetProcessProfile(ctx context.Context, profileId string) (*ProcessProfile, error) {
	q := `select * from process_profiles where id = $1`
	a := ProcessProfile{}
	if err := r.conn.GetContext(ctx, &a, q, profileId); err != nil {
		return nil, err
	}
	return &a, nil
}
