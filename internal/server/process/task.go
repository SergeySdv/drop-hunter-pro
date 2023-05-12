package process

import (
	"context"
	"crypto_scripts/internal/log"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
)

type TaskArg struct {
	Task      *v1.ProcessTask
	ProfileId string
	Rep       *repository.PGRepository
	ProcessId string
	UserId    string
}

type TaskRes struct {
	Task *v1.ProcessTask
}

type Tasker interface {
	Run(ctx context.Context, arg *TaskArg) (*TaskRes, error)
}

var executors = map[v1.TaskType]Tasker{
	v1.TaskType_StargateBridge:   &Wrap{Tasker: &StargateTask{}},
	v1.TaskType_Mock:             &Wrap{Tasker: &mockTask{}},
	v1.TaskType_Delay:            &Wrap{Tasker: &taskDelay{}},
	v1.TaskType_WithdrawExchange: &Wrap{Tasker: &WithdrawExchange{}},
}

type Wrap struct {
	Tasker Tasker
}

func (w *Wrap) Run(ctx context.Context, a *TaskArg) (task *TaskRes, err error) {

	taskId := a.Task.Id

	defer func() {
		if err != nil {
			e := err.Error()
			a.Task.Error = &e
			before := a.Task.Status
			a.Task.Status = v1.ProcessStatus_StatusError
			if err := a.Rep.UpdateProcessTask(ctx, a.Task, taskId, a.ProcessId, a.ProfileId); err != nil {
				log.Log.Error("arg.Rep.UpdateProcessTask", err)
			}

			_ = a.Rep.RecordStatusChanged(ctx, taskId, before, v1.ProcessStatus_StatusError, e)
		}

	}()

	status, err := a.Rep.GetProcessStatus(ctx, a.ProcessId)
	if err != nil {
		return nil, err
	}
	if *status == v1.ProcessStatus_StatusStop {
		return &TaskRes{Task: a.Task}, nil
	}

	task, err = w.Tasker.Run(ctx, a)

	return task, err
}
