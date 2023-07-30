package process

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

const (
	delayBeforePPTask = time.Second * 10
	maxTaskRetry      = 3
)

func getDelayBeforePPTask() time.Duration {
	if config.CFG.Env == "local" {
		return time.Millisecond * 100
	}
	return delayBeforePPTask
}

func getMaxRetry() int {
	if config.CFG.Env == "local" {
		return 0
	}
	return maxTaskRetry
}

func (d *Dispatcher) RunPP(ctx context.Context, profileId, processId, userId string) error {

	l := log.Log.With("ppId", profileId).With("pId", processId)
	l.Debug("ppId started")
	defer l.Debug("ppId finished")

	processMap, ok := d.pTable.Get(processId)
	if !ok {
		return nil
	}
	profileBusy, ok := processMap.ppTable.Get(profileId)
	if !ok {
		return nil
	}

	if profileBusy {
		return nil
	}
	processMap.ppTable.Set(profileId, true)
	defer processMap.ppTable.Set(profileId, false)

	tr := otel.Tracer("")
	pctx, span := tr.Start(ctx, "RunPP")
	span.SetAttributes(attribute.String("ppId", profileId), attribute.String("pId", processId))
	defer span.End()

	d.stat.ActiveProfiles.Inc()
	defer d.stat.ActiveProfiles.Dec()

	err := d.runPP(pctx, processId, profileId, userId, l)

	if _, err := d.resolvePPStatus(pctx, profileId, l); err != nil {
		return errors.Wrap(err, "RunPP")
	}

	return err
}

type ExecutorResult struct {
	Task *v1.ProcessTask
	Err  error
}

func (d *Dispatcher) runPP(ctx context.Context, pId processId, ppId string, userId string, l *zap.SugaredLogger) error {
	pt, _ := d.pTable.Get(pId)
	taskFinished := make(chan bool)
	defer close(taskFinished)

	pp, err := d.LoadPP(ctx, ppId)
	if err != nil {
		return errors.Wrap(err, "LoadPP")
	}

	if pp.Status() == v1.ProcessStatus_StatusDone {
		return nil
	}

	pp.SetStatus(v1.ProcessStatus_StatusRunning)

	if err := d.UpdatePPStatus(ctx, ppId, pId, v1.ProcessStatus_StatusRunning); err != nil {
		return errors.Wrap(err, "UpdatePPStatus ")
	}

	u, err := d.runner.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	retryCount := 0
	for {

		t, err := pp.GetTaskToExec()
		if errors.Is(err, ErrNoTaskToExec) {
			l.Debug("pp.GetTaskToExec: ", err)
			return nil
		}

		executor, err := task.GetTask(t.Task.TaskType)
		if err != nil {
			return err
		}

		result := make(chan ExecutorResult)

		go func() {

			d.stat.ActiveTasks.Inc()
			defer d.stat.ActiveTasks.Dec()

			execTask, execErr := executor.Run(ctx,
				&task.Input{
					L:                    l,
					Task:                 t,
					ProfileId:            pp.pp.ProfileId,
					ProcessRepository:    d.runner.processRepository,
					ProfileRepository:    d.runner.profileRepository,
					WithdrawerRepository: d.runner.withdrawerRepository,
					ProcessId:            pId,
					UserId:               userId,
					SettingsService:      d.settingsService,
					Snapshot:             d.snapshotService,
					User:                 u,
					Halper:               halp.NewHalp(d.runner.profileRepository, d.settingsService),
					PayService:           d.payService,
					Orbiter:              d.orbiterService,
				})

			result <- ExecutorResult{
				Task: execTask,
				Err:  execErr,
			}
		}()

		var executed *v1.ProcessTask
		var execErr error
		stopped := false

		go func() {
			for {
				select {
				case res := <-result:
					execErr = res.Err
					executed = res.Task
					taskFinished <- true
					return
				case signal := <-pt.signals:
					switch signal {
					case SignalStop:
						_ = executor.Stop()
						stopped = true
					}
				}
			}
		}()

		<-taskFinished

		if stopped {
			if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusStop.String(), t.Id, pId); err != nil {
				return errors.Wrap(err, "UpdateProcessTaskStatus")
			}
			return nil
		}

		if execErr != nil {

			if errors.Is(execErr, task.ErrUserHasNoBalance) {
				d.SendAlert(ctx, userId, pId, execErr, l)
				return execErr
			}

			retryCount++
			if retryCount < getMaxRetry() {
				err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRetry.String(), t.Id, pId)
				if err != nil {
					return errors.Wrap(err, "UpdateProcessTaskStatus")
				}
				pp, err = d.LoadPP(ctx, ppId)
				if err != nil {
					return errors.Wrap(err, "LoadPP")
				}
				continue
			}

			d.SendAlert(ctx, userId, pId, execErr, l)
		} else {
			retryCount = 0
			t = executed
			if err := d.r.UpdateProcessTaskStatus(ctx, t.Status.String(), t.Id, pId); err != nil {
				return err
			}
			// no time to explain
			if executed.Status == v1.ProcessStatus_StatusDone || t.Task.TaskType == v1.TaskType_Delay {
				continue
			}
		}

		status, err := d.resolvePPStatus(ctx, ppId, l)
		if err != nil {
			return errors.Wrap(err, "resolvePPStatus")
		}

		switch *status {
		case v1.ProcessStatus_StatusError,
			v1.ProcessStatus_StatusDone,
			v1.ProcessStatus_StatusStop:
			return nil
		}

		signal := <-d.sleep(ctx, pId, getDelayBeforePPTask())
		switch signal {
		case SignalStop:
			if err := d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusStop.String(), t.Id, pId); err != nil {
				return errors.Wrap(err, "UpdateProcessTaskStatus")
			}
			return nil
		}
	}
}

func (d *Dispatcher) SendAlert(ctx context.Context, userId string, pId string, execErr error, l *zap.SugaredLogger) {

	if execErr == nil {
		return
	}

	chatId, err := d.runner.userRepository.GetUserTelegramChatId(ctx, userId)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {

		} else {
			l.Warn("sending alert error", zap.Error(err))
		}
		return
	}
	if err := d.runner.tgBot.SendAlert(*chatId, pId, execErr.Error()); err != nil {
		l.Warn("SendAlert", zap.Error(err))
	}
}

func (d *Dispatcher) NotifyUserProcessStarted(ctx context.Context, userId string, pId string) {

	chatId, err := d.runner.userRepository.GetUserTelegramChatId(ctx, userId)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {

		} else {
		}
		return
	}
	if err := d.runner.tgBot.ProcessStarted(*chatId, pId); err != nil {
	}
}

func (d *Dispatcher) NotifyUserProcessFinished(ctx context.Context, userId string, pId string, l *zap.SugaredLogger) {

	chatId, err := d.runner.userRepository.GetUserTelegramChatId(ctx, userId)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {

		} else {
			l.Warn("sending alert error", zap.Error(err))
		}
		return
	}
	if err := d.runner.tgBot.ProcessFinished(*chatId, pId); err != nil {
		l.Warn("tgBot.ProcessStarted", zap.Error(err))
	}
}

func (d *Dispatcher) LoadPP(ctx context.Context, profileId string) (*PP, error) {
	ppDB, err := d.runner.processRepository.GetProcessProfileArg(ctx, profileId)
	if err != nil {
		return nil, errors.Wrap(err, "GetProcessProfileArg")
	}

	pp, err := NewPP(ppDB)
	if err != nil {
		return nil, errors.Wrap(err, "NewPP")
	}
	return pp, nil
}

func (d *Dispatcher) UpdatePPStatus(ctx context.Context, ppId, pId string, status v1.ProcessStatus) error {
	if err := d.r.UpdateProcessProfileStatus(ctx, ppId, pId, status.String()); err != nil {
		return errors.Wrap(err, "UpdateProcessProfileStatus")
	}
	if err := d.r.UpdateProcessTime(ctx, pId); err != nil {
		return errors.Wrap(err, "UpdateProcessTime")
	}
	return nil
}

func (d *Dispatcher) sleep(ctx context.Context, processId processId, sleepTime time.Duration) chan Signal {
	signals := make(chan Signal)

	go func() {
		processTable, _ := d.pTable.Get(processId)
		defer close(signals)
		timer := time.NewTimer(sleepTime)
		for {
			select {
			case <-timer.C:
				signals <- SignalWakeup
				return
			case <-ctx.Done():
				signals <- SignalTimeout
				return
			case signal := <-processTable.signals:
				signals <- signal
			}
		}
	}()

	return signals
}

func (d *Dispatcher) resolvePPStatus(ctx context.Context, profileId string, l *zap.SugaredLogger) (*v1.ProcessStatus, error) {

	pp, err := d.LoadPP(ctx, profileId)
	if err != nil {
		return nil, errors.Wrap(err, "GetProcessProfileArg")
	}

	if pp.Status() == v1.ProcessStatus_StatusStop {
		tmp := pp.Status()
		return &tmp, nil
	}

	status := resolvePPStatus(pp.TaskStatuses())
	if pp.Status() != status {
		l.Info("profile status: ", pp.Status().String(), " -> ", status.String())
	}

	if err := d.runner.processRepository.UpdateProcessProfileStatus(ctx, profileId, pp.ProcessId(), status.String()); err != nil {
		return nil, errors.Wrap(err, "UpdateProcessProfileStatus")
	}

	if err := d.runner.processRepository.UpdateProcessTime(ctx, pp.ProcessId()); err != nil {
		return nil, errors.Wrap(err, "UpdateProcessTime")
	}
	return &status, nil
}

func resolvePPStatus(statuses []v1.ProcessStatus) v1.ProcessStatus {

	done := len(statuses)
	for _, status := range statuses {

		if status == v1.ProcessStatus_StatusStop {
			return v1.ProcessStatus_StatusStop
		}

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
