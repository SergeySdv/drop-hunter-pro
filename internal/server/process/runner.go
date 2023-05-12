package process

import (
	"context"
	"crypto_scripts/internal/log"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
	"sync"

	"github.com/pkg/errors"
)

type Runner struct {
	my         *sync.Mutex
	r          *repository.PGRepository
	persistent bool
}

func NewRunner(r *repository.PGRepository, persistent bool) *Runner {
	return &Runner{
		my:         &sync.Mutex{},
		r:          r,
		persistent: persistent,
	}
}

func (r *Runner) RunProfile(ctx context.Context, profileId, userId string) error {

	l := log.Log.With("profileId", profileId).With("fn", "Runner.RunProfile")

	profileDB, err := r.r.GetProcessProfileArg(ctx, profileId)
	if err != nil {
		return errors.Wrap(err, "GetProcessProfileArg")
	}
	processId := profileDB.ProcessId

	profilePB, err := profileDB.ToPB()
	if err != nil {
		return errors.Wrap(err, "profileDB.ToPB")
	}

	profile := NewProfile(profilePB)

	if profile.Status() == v1.ProcessStatus_StatusDone {
		return nil
	}

	task, err := profile.GetTaskToExec()
	if errors.Is(err, ErrNoTaskToExec) {
		l.Debug("profile.GetTaskToExec: ", err)
		return nil
	}

	profile.p.Status = v1.ProcessStatus_StatusRunning

	profile.p, err = r.r.UpdateProcessProfile(ctx, profile.p, processId)
	if err != nil {
		return errors.Wrap(err, "UpdateProcessProfile ")
	}

	executedWithError := false

	for task != nil && !executedWithError {

		task, err = profile.GetTaskToExec()
		if errors.Is(err, ErrNoTaskToExec) {
			l.Debug("profile.GetTaskToExec: ", err)
			break
		}

		l.Debug("executors.Run: ", task.Id, "status: ", task.Status)
		executed, err := executors[task.Task.TaskType].Run(ctx,
			&TaskArg{
				Task:      task,
				ProfileId: profile.p.ProfileId,
				Rep:       r.r,
				ProcessId: processId,
				UserId:    userId,
			})
		if err != nil {
			l.Error("executors.Run ", err)
			executedWithError = true
		} else {
			task = executed.Task
		}

		statuses, err := r.r.GetProcessProfileTaskStatuses(ctx, profileId)
		if err != nil {
			return errors.Wrap(err, "GetProcessProfileTaskStatuses ")
		}

		status := GetProfileStatus(statuses)

		profile.p.Status = status

		// обновление profile
		profile.p, err = r.r.UpdateProcessProfile(ctx, profile.p, processId)
		if err != nil {
			return errors.Wrap(err, "UpdateProcessProfile ")
		}
		_, err = r.ResolveProcessStatus(ctx, processId)
		if err != nil {
			return errors.Wrap(err, "ResolveProcessStatus ")
		}
	}

	return nil
}

func GetProfileStatus(statuses []v1.ProcessStatus) v1.ProcessStatus {

	done := len(statuses)
	for _, status := range statuses {
		if status == v1.ProcessStatus_StatusError {
			return v1.ProcessStatus_StatusError
		}
		if status == v1.ProcessStatus_StatusDone {
			done--
		}
	}

	if done == 0 {
		return v1.ProcessStatus_StatusDone
	}

	return v1.ProcessStatus_StatusRunning
}

func (d *Runner) ResolveProcessStatus(ctx context.Context, processId string) (*v1.ProcessStatus, error) {

	ps, err := d.r.GetProcessStatus(ctx, processId)
	if err != nil {
		return nil, errors.Wrap(err, "GetProcessStatus")
	}
	if *ps == v1.ProcessStatus_StatusStop {
		return ps, nil
	}

	statuses, err := d.r.GetProcessProfileStatuses(ctx, processId)
	if err != nil {
		return nil, err
	}

	status := GetProcessStatus(statuses)

	err = d.r.UpdateProcessStatus(ctx, &repository.UpdateProcess{Id: processId, Status: status.String()})
	if err != nil {
		return nil, err
	}
	return &status, nil
}
