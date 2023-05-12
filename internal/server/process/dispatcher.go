package process

import (
	"context"
	"crypto_scripts/internal/log"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"

	"github.com/pkg/errors"
)

type processId = string
type tableArg struct {
	profiles *Map[string, bool]
	order    []string
}
type Dispatcher struct {
	processMap *Map[processId, *tableArg]
	r          *repository.PGRepository
	runner     *Runner
	queue      chan processId
	eror       chan error
}

func NewDispatcher(r *repository.PGRepository, runner *Runner) *Dispatcher {
	d := &Dispatcher{
		processMap: NewMap[processId, *tableArg](),
		r:          r,
		runner:     runner,
		queue:      make(chan processId, 1000),
		eror:       make(chan error, 1000),
	}

	return d

}

func (d *Dispatcher) RetryProcessProfile(ctx context.Context, processId, taskId string) error {

	status, err := d.r.GetTaskStatus(ctx, taskId)
	if err != nil {
		return err
	}

	if *status != v1.ProcessStatus_StatusError {
		return errors.New("forbidden action")
	}

	if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRetry.String(), taskId, processId); err != nil {
		return err
	}

	go d.StartProcess(processId)

	return nil
}

func (d *Dispatcher) ResumeProcess(processId string) {

	if err := d.r.UpdateProcessStatus(context.Background(), &repository.UpdateProcess{
		Id:     processId,
		Status: v1.ProcessStatus_StatusRunning.String(),
	}); err != nil {
		return
	}

	d.queue <- processId
}

func (d *Dispatcher) StartProcess(processId string) {
	d.queue <- processId
}

func (d *Dispatcher) KillProcess(processId string) error {

	status, err := d.r.GetProcessStatus(context.Background(), processId)
	if err != nil {
		return err
	}

	if *status == v1.ProcessStatus_StatusStop ||
		*status == v1.ProcessStatus_StatusDone ||
		*status == v1.ProcessStatus_StatusError {
		if err := d.StopProcess(processId); err != nil {
			return err
		}

		if err := d.r.DeleteProcess(context.Background(), processId); err != nil {
			return err
		}
	}

	return errors.New("can't kill process in that status: " + (*status).String())
}

func (d *Dispatcher) StopProcess(processId string) error {
	return d.r.UpdateProcessStatus(context.Background(), &repository.UpdateProcess{
		Id:     processId,
		Status: v1.ProcessStatus_StatusStop.String(),
	})
}

func (d *Dispatcher) Run(ctx context.Context) error {

	go func() {
		processIds, err := d.r.ListProcessIdsByStatus(ctx, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning)
		if err != nil {
			d.eror <- err
		}

		for _, processId := range processIds {
			d.queue <- processId
		}
	}()

	for processId := range d.queue {

		l := log.Log.With("processId", processId).With("fn", "Dispatcher.Run")
		l.Debug("got from queue")
		// todo: фильтровать статусы

		status, err := d.r.GetProcessStatus(ctx, processId)
		if err != nil {
			l.Error("GetProcessStatus")
			return err
		}
		if *status == v1.ProcessStatus_StatusStop {
			return nil
		}

		_, exist := d.processMap.Get(processId)
		if exist {
			l.Debug("already processing")
			continue
		}

		profileIds, err := d.r.GetProcessProfileIds(ctx, processId)
		if err != nil {
			l.Error("GetProcessProfileIds")
			return err
		}

		profileMap := &tableArg{
			profiles: NewMap[string, bool](),
			order:    profileIds,
		}
		for _, profileId := range profileIds {
			profileMap.profiles.Set(profileId, false)
		}

		d.processMap.Set(processId, profileMap)

		go d.RunProcess(ctx, processId)
	}

	return nil
}

func (d *Dispatcher) RunProcess(ctx context.Context, processId string) error {

	l := log.Log.With("processId", processId).With("fn", "Dispatcher.RunProcess")

	err := d.r.UpdateProcessStatus(ctx, &repository.UpdateProcess{Id: processId, Status: v1.ProcessStatus_StatusRunning.String()})
	if err != nil {
		l.Error("UpdateProcessStatus", err)
		return err
	}

	profMap, _ := d.processMap.Get(processId)

	for _, profileId := range profMap.order {

		profMap.profiles.Set(processId, true)

		status, err := d.r.GetProcessStatus(ctx, processId)
		if err != nil {
			return errors.Wrap(err, "GetProcessStatus")
		}

		userId, err := d.r.GetProcessUser(ctx, processId)
		if err != nil {
			return errors.Wrap(err, "GetProcessUser")
		}

		if *status == v1.ProcessStatus_StatusStop {
			return nil
		}

		err = d.runner.RunProfile(ctx, profileId, *userId)
		if err != nil {
			l.Error("runner.RunProfile", err)
			profMap.profiles.Set(processId, false)
			return err
		}

		status, err = d.ResolveProcessStatus(ctx, processId)
		if err != nil {
			l.Error("ResolveProcessStatus", err)
			profMap.profiles.Set(processId, false)
			return err
		}

		profMap.profiles.Set(processId, false)

		switch *status {
		case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
			d.processMap.Remove(processId)
			return nil
		}
	}

	return nil
}
func (d *Dispatcher) ResolveProcessStatus(ctx context.Context, processId string) (*v1.ProcessStatus, error) {

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

func GetProcessStatus(profileStatuses []v1.ProcessStatus) v1.ProcessStatus {

	done := len(profileStatuses)
	for _, status := range profileStatuses {
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
