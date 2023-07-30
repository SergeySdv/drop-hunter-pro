package task

import (
	"context"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mockTask struct {
	cancel func()
}

func (t *mockTask) Stop() error {
	t.cancel()
	return nil
}

func (t *mockTask) Reset(ctx context.Context, a *Input) error {
	return nil
}

func (t *mockTask) Run(ctx context.Context, arg *Input) (*v1.ProcessTask, error) {
	_, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	arg.Task.Status = v1.ProcessStatus_StatusRunning

	tt := arg.Task
	tt.Status = v1.ProcessStatus_StatusDone
	tt.FinishedAt = timestamppb.Now()

	return tt, nil
}
