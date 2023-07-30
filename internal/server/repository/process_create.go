package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (r *pgRepository) CreateProcess(ctx context.Context, req *ProcessArg) error {

	tx, err := r.conn.BeginTxx(ctx, nil)
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
