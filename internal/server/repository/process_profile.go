package repository

import (
	"context"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
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
	Label     string `db:"label"`
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
		Id:           a.Id,
		ProfileId:    a.ProfileId,
		Weight:       a.Weight,
		Tasks:        nil,
		ProfileLabel: a.Label,
		Status:       v1.ProcessStatus(v1.ProcessStatus_value[a.Status]),
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

func (r *pgRepository) UpdateProcessProfileStatus(ctx context.Context, profileId, processId string, status string) error {
	updatedAt := time.Now()

	tx, err := r.conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	q := `update process_profiles set status = $1  where id = $2`
	if _, err := r.conn.ExecContext(ctx, q, status, profileId); err != nil {
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

func (r *pgRepository) GetProcessProfileIds(ctx context.Context, processId string) ([]string, error) {
	q := `select id from process_profiles where process_id = $1 order by weight asc`
	a := make([]string, 0)
	if err := r.conn.SelectContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *pgRepository) GetProcessProfileStatus(ctx context.Context, profileId string) (*v1.ProcessStatus, error) {
	q := `select status from process_profiles where id = $1`
	var s string
	if err := r.conn.GetContext(ctx, &s, q, profileId); err != nil {
		return nil, err
	}

	t := v1.ProcessStatus(v1.ProcessStatus_value[s])
	return &t, nil
}

func (r *pgRepository) GetProcessProfileStatuses(ctx context.Context, processId string) ([]v1.ProcessStatus, error) {
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

func (r *pgRepository) GetProcessProfileArg(ctx context.Context, profileId string) (*ProcessProfileArg, error) {

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

func (r *pgRepository) GetProcessProfile(ctx context.Context, profileId string) (*ProcessProfile, error) {
	q := `select pp.*, p.label as label  from process_profiles as pp join profiles as p on pp.profile_id = p.id where pp.id = $1`
	a := ProcessProfile{}
	if err := r.conn.GetContext(ctx, &a, q, profileId); err != nil {
		return nil, err
	}
	return &a, nil
}
