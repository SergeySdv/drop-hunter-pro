package task

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type taskDelay struct {
	cancel func()
}

func (t *taskDelay) Stop() error {
	t.cancel()
	return nil
}

func (t *taskDelay) Reset(ctx context.Context, a *Input) error {
	task := a.Task
	l, ok := task.Task.Task.(*v1.Task_DelayTask)
	if !ok {
		return errors.New("a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}
	p := l.DelayTask

	p.WaitFor = nil

	if err := a.UpdateTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (t *taskDelay) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	task := a.Task
	l, ok := task.Task.Task.(*v1.Task_DelayTask)
	if !ok {
		return nil, errors.New("a.Task.Task.Task.(*v1.Task_StargateBridgeTask) call an ambulance!")
	}
	p := l.DelayTask

	switch task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:
		task.Status = v1.ProcessStatus_StatusRunning
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
			rd := int64(duration.Seconds())
			p.Duration = rd
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}

	case v1.ProcessStatus_StatusRunning:
		waitFor := l.DelayTask.WaitFor.AsTime().Unix()
		now := time.Now().Unix()

		if waitFor-now <= 0 {
			task.Status = v1.ProcessStatus_StatusDone
			task.FinishedAt = timestamppb.Now()

			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, err
			}
			return task, nil
		}
		duration := waitFor - now
		d := time.Duration(duration)
		d *= time.Second

		taskContext, cancel := context.WithTimeout(ctx, d)
		defer cancel()

		t.cancel = cancel

		timer := time.NewTimer(d)
		defer timer.Stop()
		select {
		case <-timer.C:
			task.Status = v1.ProcessStatus_StatusDone
			task.FinishedAt = timestamppb.Now()

			if err := a.UpdateTask(ctx, task); err != nil {
				return nil, err
			}
		case <-taskContext.Done():
			return task, nil
		}

	}

	return a.Task, nil
}
