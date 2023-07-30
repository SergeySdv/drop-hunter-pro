package repository

import (
	"context"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	Id        string    `db:"id"`
	Code      string    `db:"code"`
	Network   string    `db:"network"`
	Url       string    `db:"url"`
	TaskId    string    `db:"task_id"`
	ProfileId string    `db:"profile_id"`
	ProcessId string    `db:"process_id"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

var TransactionCols = []string{
	"id",
	"code",
	"task_id",
	"profile_id",
	"process_id",
	"user_id",
	"created_at",
	"network",
	"url",
}

var (
	trfh = NewHelper(TransactionCols)
)

func (a *Transaction) ToPB() *v1.Transaction {

	p := &v1.Transaction{
		Id:        a.Id,
		Code:      a.Code,
		TaskId:    a.TaskId,
		ProfileId: a.ProfileId,
		ProcessId: a.ProcessId,
		UserId:    a.UserId,
		CreatedAt: timestamppb.New(a.CreatedAt),
		Network:   v1.Network(v1.Network_value[a.Network]),
		Url:       a.Url,
	}

	return p
}

func (r *pgRepository) TransactionAdd(ctx context.Context, req *Transaction) error {
	q := "insert into transactions (" + trfh.Cols() + ") values (" + trfh.ColsColon() + ") on conflict (id) do nothing"

	_, err := r.conn.NamedExecContext(ctx, q, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *pgRepository) TransactionsByTaskId(ctx context.Context, userId, taskId string) ([]Transaction, error) {
	q := "select " + trfh.Cols() + " from transactions where user_id = $1 and task_id = $2 order by created_at desc"

	out := make([]Transaction, 0)
	if err := r.conn.SelectContext(ctx, &out, q, userId, taskId); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *pgRepository) TransactionsByProfileId(ctx context.Context, userId, profileId string) ([]Transaction, error) {
	q := "select " + trfh.Cols() + " from transactions where user_id = $1 and profile_id = $2 order by created_at desc limit 5"

	out := make([]Transaction, 0)
	if err := r.conn.SelectContext(ctx, &out, q, userId, profileId); err != nil {
		return nil, err
	}
	return out, nil
}
