package process

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type mockTask struct {
}

func (t *mockTask) Run(ctx context.Context, arg *TaskArg) (*TaskRes, error) {

	arg.Task.Status = v1.ProcessStatus_StatusRunning

	tt := arg.Task
	tt.Status = v1.ProcessStatus_StatusDone
	tt.FinishedAt = timestamppb.Now()

	return &TaskRes{
		Task: tt,
	}, nil
}
