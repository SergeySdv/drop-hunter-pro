package process

import (
	"context"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type taskDelay struct {
}

func (t *taskDelay) Run(ctx context.Context, a *TaskArg) (*TaskRes, error) {

	task := a.Task
	l, ok := task.Task.Task.(*v1.Task_DelayTask)
	if !ok {
		return nil, errors.New("a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}
	p := l.DelayTask

	switch task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return &TaskRes{Task: task}, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:
		before := task.Status
		after := v1.ProcessStatus_StatusRunning
		task.Status = after
		if p.WaitFor == nil {
			var duration time.Duration
			if p.Random {
				min := time.Duration(*p.MinRandom) * time.Second
				max := time.Duration(*p.MaxRandom) * time.Second
				duration = lib.RandDurationRange(min, max)
			} else {
				duration = time.Duration(p.Duration) * time.Second
			}

			p.WaitFor = timestamppb.New(time.Now().Add(duration))
			rd := duration.String()
			p.RandomDuration = &rd
		}
		if err := a.Rep.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, err
		}
		_ = a.Rep.RecordStatusChanged(ctx, task.Id, before, after)

	case v1.ProcessStatus_StatusRunning:
		if time.Now().After(l.DelayTask.WaitFor.AsTime()) {
			task.Status = v1.ProcessStatus_StatusDone
			task.FinishedAt = timestamppb.Now()

			if err := a.Rep.UpdateProcessTaskStatus(ctx, task.Status.String(), task.Id, a.ProcessId); err != nil {
				return nil, err
			}
			_ = a.Rep.RecordStatusChanged(ctx, task.Id, v1.ProcessStatus_StatusRunning, v1.ProcessStatus_StatusDone)
		}
	}

	return &TaskRes{
		Task: a.Task,
	}, nil
}
