package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProcessTaskHistory struct {
	Id         string       `db:"id"`
	TaskId     string       `db:"task_id"`
	StartedAt  time.Time    `db:"started_at"`
	FinishedAt sql.NullTime `db:"finished_at"`

	StartStatus  string         `db:"start_status"`
	FinishStatus sql.NullString `db:"finish_status"`
	Msg          sql.NullString `db:"msg"`
}

type Record struct {
	record *ProcessTaskHistory
	r      *pgRepository
}

func (r *pgRepository) GetProcessTaskHistory(ctx context.Context, taskId string) ([]*v1.ProcessTaskHistoryRecord, error) {

	q := `select * from process_task_history where  task_id = $1 order by started_at desc limit 15`

	records := make([]ProcessTaskHistory, 0)
	if err := r.conn.SelectContext(ctx, &records, q, taskId); err != nil {
		return nil, err
	}

	tmp := make([]*v1.ProcessTaskHistoryRecord, 0)
	for i := range records {
		record := &records[i]
		var rect = v1.ProcessTaskHistoryRecord{
			Id:           record.Id,
			TaskId:       record.TaskId,
			StartedAt:    timestamppb.New(record.StartedAt),
			FinishedAt:   nil,
			StartStatus:  v1.ProcessStatus(v1.ProcessStatus_value[record.StartStatus]),
			FinishStatus: nil,
			Msg:          nil,
		}

		if record.FinishStatus.Valid {
			t := v1.ProcessStatus(v1.ProcessStatus_value[record.FinishStatus.String])
			rect.FinishStatus = &t
		}
		if record.FinishedAt.Valid {
			rect.FinishedAt = timestamppb.New(record.FinishedAt.Time)
		}

		if record.Msg.Valid {
			rect.Msg = &record.Msg.String
		}
		tmp = append(tmp, &rect)
	}
	return tmp, nil
}

func (r *pgRepository) RecordStatusChanged(ctx context.Context, taskId string, before, after v1.ProcessStatus, msg ...string) error {
	rec, err := r.StartRecord(ctx, taskId, before.String())
	if err != nil {
		return err
	}
	return rec.FinishRecord(ctx, after.String(), strings.Join(msg, "\n"))
}

func (r *pgRepository) StartRecord(ctx context.Context, taskId, startStatus string) (*Record, error) {

	record := ProcessTaskHistory{
		Id:           uuid.New().String(),
		TaskId:       taskId,
		StartedAt:    time.Now(),
		FinishedAt:   sql.NullTime{},
		StartStatus:  startStatus,
		FinishStatus: sql.NullString{},
		Msg:          sql.NullString{},
	}

	q := `insert into process_task_history
    (id, task_id, started_at, start_status) values (:id, :task_id, :started_at, :start_status) `
	if _, err := r.conn.NamedExecContext(ctx, q, &record); err != nil {
		return nil, err
	}
	return &Record{
		record: &record,
		r:      r,
	}, nil
}

func (r *Record) FinishRecord(ctx context.Context, finishStatus, msg string) error {

	q := `update process_task_history set
				finished_at = :started_at,
				finish_status = :finish_status,
				msg = :msg
					where id = :id `

	record := r.record
	record.FinishedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	record.FinishStatus = sql.NullString{
		String: finishStatus,
		Valid:  true,
	}

	record.Msg = sql.NullString{
		String: msg,
		Valid:  true,
	}

	if _, err := r.r.conn.NamedExecContext(ctx, q, record); err != nil {
		return err
	}
	return nil
}
