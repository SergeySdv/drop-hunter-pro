package repository

import (
	"context"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

func (r *pgRepository) GetProcessArg(ctx context.Context, req *v1.GetProcessRequest, userId string) (*v1.GetProcessResponse, error) {

	var p ProcessArg

	process, err := r.GetProcess(ctx, req.Id, userId)
	if err != nil {
		return nil, err
	}
	p.Process = *process

	progress, err := r.GetProcessProgress(ctx, p.Id)
	if err != nil {
		return nil, err
	}
	p.Process.Progress = *progress

	profiles, err := r.GetProcessProfiles(ctx, process.Id)
	if err != nil {
		return nil, err
	}

	prfs := make([]ProcessProfileArg, 0)
	for i := range profiles {
		profile := &profiles[i]
		tasks, err := r.GetProcessTasks(ctx, profile.Id)
		if err != nil {
			return nil, err
		}
		prfs = append(prfs, ProcessProfileArg{
			ProcessProfile: *profile,
			Tasks:          tasks,
		})
	}

	p.Profiles = prfs

	flow, err := r.GetFlow(ctx, &v1.GetFlowRequest{
		Id: p.Process.FlowId,
	})
	if err != nil {
		return nil, err
	}

	pb, err := p.ToPB(flow.Flow)
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessResponse{
		Process: pb,
	}, nil
}

func (r *pgRepository) GetProcess(ctx context.Context, id, user_id string) (*Process, error) {
	q := `select * from process where id = $1 and user_id = $2`
	var a Process
	if err := r.conn.GetContext(ctx, &a, q, id, user_id); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) GetProcessProfiles(ctx context.Context, processId string) ([]ProcessProfile, error) {
	q := `select pp.*, cast(p.num as text) as label from process_profiles as pp join profiles as p on pp.profile_id = p.id where process_id = $1 order by weight asc`
	a := make([]ProcessProfile, 0)
	if err := r.conn.SelectContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *pgRepository) GetProcessTasks(ctx context.Context, profileId string) ([]ProcessTask, error) {
	q := `select * from process_tasks where profile_id = $1 order by weight asc`
	a := make([]ProcessTask, 0)
	if err := r.conn.SelectContext(ctx, &a, q, profileId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *pgRepository) GetProcessTask(ctx context.Context, taskId string) (*ProcessTask, error) {
	q := `select * from process_tasks where id = $1 order by weight asc`
	var a ProcessTask
	if err := r.conn.GetContext(ctx, &a, q, taskId); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgRepository) GetProcessTasksStatus(ctx context.Context, taskId string) (*string, error) {
	q := `select status from process_tasks where id = $1`
	var a string
	if err := r.conn.GetContext(ctx, &a, q, taskId); err != nil {
		return nil, err
	}
	return &a, nil
}
