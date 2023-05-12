package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
)

func (r *PGRepository) GetProcessArg(ctx context.Context, req *v1.GetProcessRequest) (*v1.GetProcessResponse, error) {

	var p ProcessArg

	process, err := r.GetProcess(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	p.Process = *process

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

	pb, err := p.ToPB(flow.Flow.GetLabel())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessResponse{
		Process: pb,
	}, nil
}

func (r *PGRepository) GetProcess(ctx context.Context, id string) (*Process, error) {
	q := `select * from process where id = $1`
	var a Process
	if err := r.conn.GetContext(ctx, &a, q, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *PGRepository) GetProcessProfiles(ctx context.Context, processId string) ([]ProcessProfile, error) {
	q := `select * from process_profiles where process_id = $1 order by weight asc`
	a := make([]ProcessProfile, 0)
	if err := r.conn.SelectContext(ctx, &a, q, processId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *PGRepository) GetProcessTasks(ctx context.Context, profileId string) ([]ProcessTask, error) {
	q := `select * from process_tasks where profile_id = $1 order by weight asc`
	a := make([]ProcessTask, 0)
	if err := r.conn.SelectContext(ctx, &a, q, profileId); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *PGRepository) GetProcessTasksStatus(ctx context.Context, taskId string) (*string, error) {
	q := `select status from process_tasks where id = $1`
	var a string
	if err := r.conn.GetContext(ctx, &a, q, taskId); err != nil {
		return nil, err
	}
	return &a, nil
}
