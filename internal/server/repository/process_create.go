package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/user"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *PGRepository) CreateProcess(ctx context.Context, req *v1.CreateProcessRequest) (*v1.CreateProcessResponse, error) {

	flow, err := r.GetFlow(ctx, &v1.GetFlowRequest{
		Id: req.FlowId,
	})
	if err != nil {
		return nil, err
	}

	profiles := make([]*v1.ProcessProfile, 0)

	for profileWeight, profileId := range req.ProfileIds {

		processTasks := make([]*v1.ProcessTask, 0)

		for _, t := range flow.Flow.Tasks {
			processTasks = append(processTasks, &v1.ProcessTask{
				Id:           uuid.New().String(),
				Task:         t,
				Status:       v1.ProcessStatus_StatusReady,
				Transactions: []string{},
				FinishedAt:   nil,
				Error:        nil,
				StartedAt:    nil,
			})
		}

		profiles = append(profiles, &v1.ProcessProfile{
			ProfileId: profileId,
			Id:        uuid.New().String(),
			Weight:    int64(profileWeight),
			Tasks:     processTasks,
			Status:    v1.ProcessStatus_StatusReady,
		})
	}

	userId := user.GetUserId(ctx)
	var pb = &v1.Process{
		Id:         uuid.New().String(),
		Status:     v1.ProcessStatus_StatusStop,
		Profiles:   profiles,
		FlowId:     req.FlowId,
		CreatedAt:  timestamppb.Now(),
		UpdatedAt:  timestamppb.Now(),
		FinishedAt: nil,
		StartedAt:  nil,
	}

	a := &ProcessArg{}
	if err := a.FromPB(pb, userId); err != nil {
		return nil, err
	}

	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	if err := CreateProcess(ctx, r.conn, a); err != nil {
		return nil, err
	}

	res, err := r.GetProcessArg(ctx, &v1.GetProcessRequest{Id: a.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CreateProcessResponse{
		Process: res.Process,
	}, nil
}

func CreateProcess(ctx context.Context, conn *sqlx.DB, req *ProcessArg) error {

	tx, err := conn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := createProcess(ctx, tx, req); err != nil {
		return err
	}

	for i := range req.Profiles {
		profile := &req.Profiles[i]

		if err := createProcessProfile(ctx, tx, profile); err != nil {
			return err
		}

		for j := range profile.Tasks {
			task := &profile.Tasks[j]

			if err := createProcessTask(ctx, tx, task); err != nil {
				return err
			}
		}

	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func createProcess(ctx context.Context, conn *sqlx.Tx, req *ProcessArg) error {
	q := `insert into process (id, status, payload, user_id, flow_id, updated_at, created_at) values 
      (:id, :status, :payload, :user_id, :flow_id, :updated_at, :created_at)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func createProcessProfile(ctx context.Context, conn *sqlx.Tx, req *ProcessProfileArg) error {
	q := `insert into process_profiles (id, weight, process_id, profile_id, status) values 
      (:id, :weight, :process_id, :profile_id, :status)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}

func createProcessTask(ctx context.Context, conn *sqlx.Tx, req *ProcessTask) error {
	q := `insert into process_tasks (id, weight, process_id, profile_id, status, payload) values 
      (:id, :weight, :process_id, :profile_id, :status, :payload)                                                                 `
	if _, err := conn.NamedExecContext(ctx, q, req); err != nil {
		return err
	}
	return nil
}
