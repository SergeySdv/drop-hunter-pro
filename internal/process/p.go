package process

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

func (d *Dispatcher) RunP(ctx context.Context, processId string) error {

	l := log.Log.With("processId", processId).With("fn", "RunP")

	err := d.r.UpdateProcessStatus(ctx, &repository.UpdateProcess{Id: processId, Status: v1.ProcessStatus_StatusRunning.String()})
	if err != nil {
		l.Error("UpdateProcessStatus", err)
		return err
	}

	l.Debug("process running")

	ppTable, running := d.pTable.Get(processId)
	if !running {
		return errors.New("process not running")
	}

	d.stat.ActiveProcesses.Inc()
	defer d.stat.ActiveProcesses.Dec()

	tr := otel.Tracer("")
	pctx, span := tr.Start(ctx, "RunP")
	span.SetAttributes(attribute.String("pId", processId))
	defer span.End()

	for _, ppId := range ppTable.ppOrder {

		ppStatus, err := d.r.GetProcessProfileStatus(pctx, ppId)
		if err != nil {
			return err
		}
		switch *ppStatus {
		case v1.ProcessStatus_StatusDone:
			continue
		}

		userId, err := d.r.GetProcessUser(pctx, processId)
		if err != nil {
			return errors.Wrap(err, "GetProcessUser")
		}

		err = d.RunPP(pctx, ppId, processId, *userId)
		if err != nil {
			l.Error("runner.RunProfile", err)
			return err
		}

		status, err := d.ResolvePStatus(pctx, processId, l)
		if err != nil {
			l.Error("ResolvePStatus", err)
			return err
		}
		if *status == v1.ProcessStatus_StatusDone {
			d.NotifyUserProcessFinished(ctx, *userId, processId, l)
		}

		switch *status {
		case v1.ProcessStatus_StatusError,
			v1.ProcessStatus_StatusDone,
			v1.ProcessStatus_StatusStop:
			return nil
		}

		ppStatus, err = d.r.GetProcessProfileStatus(pctx, ppId)
		if err != nil {
			return err
		}
		switch *ppStatus {
		case v1.ProcessStatus_StatusDone:
			continue
		default:
			return nil
		}
	}

	return nil
}
func (d *Dispatcher) StartProcess(processId string) {

	_, running := d.pTable.Get(processId)
	if running {
		return
	}

	d.queue <- processId
}
func (d *Dispatcher) KillProcess(ctx context.Context, processId string) error {

	if err := d.StopProcess(ctx, processId); err != nil {
		return err
	}

	if err := d.r.DeleteProcess(context.Background(), processId); err != nil {
		return err
	}

	return nil
}
func (d *Dispatcher) StopProcess(ctx context.Context, processId string) error {

	l := log.Log.With("processId", processId).With("fn", "Dispatcher.StopProcess")
	l.Debug("signal received")
	defer l.Debug("process stopped")
	processTable, running := d.pTable.Get(processId)
	if running {
		processTable.signals <- SignalStop
	} else {
		d.pTable.Set(processId, nil)
		defer d.pTable.Remove(processId)
		return d.r.UpdateProcessStatus(ctx, &repository.UpdateProcess{
			Id:     processId,
			Status: v1.ProcessStatus_StatusStop.String(),
		})
	}
	ticker := time.NewTicker(time.Second)
	for {
		status, err := d.r.GetProcessStatus(ctx, processId)
		if err != nil {
			return err
		}

		if *status == v1.ProcessStatus_StatusStop {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}

	return nil
}
func (d *Dispatcher) ResolvePStatus(ctx context.Context, processId string, l *zap.SugaredLogger) (*v1.ProcessStatus, error) {

	ps, err := d.r.GetProcessStatus(ctx, processId)
	if err != nil {
		return nil, errors.Wrap(err, "GetProcessStatus")
	}

	statuses, err := d.r.GetProcessProfileStatuses(ctx, processId)
	if err != nil {
		return nil, err
	}

	status := GetProcessStatus(statuses)

	if *ps == status {
		return ps, nil
	}

	l.Info("process status changed: ", ps.String(), " -> ", status.String())

	err = d.r.UpdateProcessStatus(ctx, &repository.UpdateProcess{Id: processId, Status: status.String()})
	if err != nil {
		return nil, errors.Wrap(err, "UpdateProcessStatus")
	}

	return &status, nil
}
func (d *Dispatcher) RetryProcess(ctx context.Context, processId string) error {

	if err := d.StopProcess(ctx, processId); err != nil {
		return err
	}

	user, err := d.r.GetProcessUser(ctx, processId)
	if err != nil {
		return err
	}

	p, err := d.r.GetProcessArg(ctx, &v1.GetProcessRequest{
		Id: processId,
	}, *user)
	if err != nil {
		return err
	}

	for _, profile := range p.Process.Profiles {
		if profile.Status != v1.ProcessStatus_StatusError && profile.Status != v1.ProcessStatus_StatusRetry {
			continue
		}

		for _, task := range profile.Tasks {
			if task.Status != v1.ProcessStatus_StatusError && task.Status != v1.ProcessStatus_StatusRetry {
				continue
			}

			if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRetry.String(), task.Id, processId); err != nil {
				return err
			}
			if err := d.r.UpdateProcessProfileStatus(ctx, task.Id, processId, v1.ProcessStatus_StatusReady.String()); err != nil {
				return err
			}
			go d.StartProcess(processId)
			return nil
		}
	}

	return nil
}
func (d *Dispatcher) ResumeProcess(ctx context.Context, processId string) error {

	userId, err := d.r.GetProcessUser(ctx, processId)
	if err != nil {
		return err
	}

	process, err := d.r.GetProcessArg(ctx, &v1.GetProcessRequest{
		Id: processId,
	}, *userId)
	if err != nil {
		return err
	}

	processStarted := false
	for pi, profile := range process.Process.Profiles {

		for ti, task := range profile.Tasks {
			if ti == 0 && pi == 0 {
				processStarted = task.Status == v1.ProcessStatus_StatusDone || task.Status == v1.ProcessStatus_StatusRunning
			}
			if task.Status == v1.ProcessStatus_StatusStop {
				if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRunning.String(), task.Id, processId); err != nil {
					return err
				}
			}
		}

		if profile.Status == v1.ProcessStatus_StatusStop {
			if err := d.r.UpdateProcessProfileStatus(ctx, profile.ProfileId, processId, v1.ProcessStatus_StatusRunning.String()); err != nil {
				return err
			}
		}
	}

	if !processStarted {
		d.NotifyUserProcessStarted(ctx, *userId, processId)
	}

	if err := d.r.UpdateProcessStatus(context.Background(), &repository.UpdateProcess{
		Id:     processId,
		Status: v1.ProcessStatus_StatusRunning.String(),
	}); err != nil {
		return err
	}

	d.queue <- processId

	return nil
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
		if status == v1.ProcessStatus_StatusStop {
			return v1.ProcessStatus_StatusStop
		}
	}

	if done == 0 {
		return v1.ProcessStatus_StatusDone
	}

	return v1.ProcessStatus_StatusRunning
}
