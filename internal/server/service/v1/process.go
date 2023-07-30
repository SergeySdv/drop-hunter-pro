package v1

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProcessService struct {
	v1.UnimplementedProcessServiceServer
	processRepository repository.ProcessRepository
	flowRepository    repository.FlowRepository
	dispatcher        *process.Dispatcher
	settingsService   *settings.Service
}

func NewProcessService(
	processRepository repository.ProcessRepository,
	dispatcher *process.Dispatcher,
	flowRepository repository.FlowRepository,
	settingsService *settings.Service,
) *ProcessService {
	return &ProcessService{
		processRepository: processRepository,
		dispatcher:        dispatcher,
		flowRepository:    flowRepository,
		settingsService:   settingsService,
	}
}

func (s *ProcessService) EnableAutoRetry(ctx context.Context, req *v1.EnableAutoRetryRequest) (*v1.EnableAutoRetryResponse, error) {
	err := s.processRepository.UpdateProcessAutoRetry(ctx, req.ProcessId, true)
	if err != nil {
		return nil, err
	}
	return &v1.EnableAutoRetryResponse{}, nil
}
func (s *ProcessService) DisableAutoRetry(ctx context.Context, req *v1.DisableAutoRetryRequest) (*v1.DisableAutoRetryResponse, error) {
	err := s.processRepository.UpdateProcessAutoRetry(ctx, req.ProcessId, false)
	if err != nil {
		return nil, err
	}
	return &v1.DisableAutoRetryResponse{}, nil
}
func (s *ProcessService) SkipProcessTask(ctx context.Context, req *v1.SkipProcessTaskRequest) (*v1.SkipProcessTaskResponse, error) {
	if err := s.dispatcher.SkipTask(ctx, req.ProcessId, req.ProfileId, req.TaskId); err != nil {
		return nil, err
	}
	return &v1.SkipProcessTaskResponse{}, nil
}
func (s *ProcessService) CreateProcess(ctx context.Context, req *v1.CreateProcessRequest) (*v1.CreateProcessResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	flow, err := s.flowRepository.GetFlow(ctx, &v1.GetFlowRequest{
		Id: req.FlowId,
	})
	if err != nil {
		return nil, err
	}

	profiles := make([]*v1.ProcessProfile, 0)

	for profileWeight, profileId := range req.ProfileIds {

		processTasks := make([]*v1.ProcessTask, 0)

		for i := range flow.Flow.Tasks {
			t := flow.Flow.Tasks[i]

			marshal, err := task.GetTaskDesc(t)
			if err != nil {
				return nil, errors.Wrap(err, "lib.Marshal")
			}
			description := string(marshal)
			t.Description = description

			processTasks = append(processTasks, &v1.ProcessTask{
				Id:           uuid.New().String(),
				Task:         t,
				Status:       v1.ProcessStatus_StatusReady,
				Transactions: []string{},
				FinishedAt:   nil,
				Error:        nil,
				StartedAt:    nil,
			})
		}

		profiles = append(profiles, &v1.ProcessProfile{
			ProfileId: profileId,
			Id:        uuid.New().String(),
			Weight:    int64(profileWeight),
			Tasks:     processTasks,
			Status:    v1.ProcessStatus_StatusReady,
		})
	}

	var pb = &v1.Process{
		Id:         uuid.New().String(),
		Status:     v1.ProcessStatus_StatusStop,
		Profiles:   profiles,
		FlowId:     req.FlowId,
		CreatedAt:  timestamppb.Now(),
		UpdatedAt:  timestamppb.Now(),
		FinishedAt: nil,
		StartedAt:  nil,
	}

	a := &repository.ProcessArg{}
	if err := a.FromPB(pb, userId); err != nil {
		return nil, err
	}

	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	if err := s.processRepository.CreateProcess(ctx, a); err != nil {
		return nil, err
	}

	res, err := s.processRepository.GetProcessArg(ctx, &v1.GetProcessRequest{Id: a.Id}, userId)
	if err != nil {
		return nil, err
	}

	return &v1.CreateProcessResponse{
		Process: res.Process,
	}, nil
}
func (s *ProcessService) GetProcess(ctx context.Context, req *v1.GetProcessRequest) (*v1.GetProcessResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	return s.processRepository.GetProcessArg(ctx, req, userId)
}
func (s *ProcessService) ListProcess(ctx context.Context, req *v1.ListProcessRequest) (*v1.ListProcessResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	statuses := make([]string, len(req.Statuses))
	for _, s := range req.Statuses {
		statuses = append(statuses, s.String())
	}

	processes, err := s.processRepository.ListProcessByUser(ctx, userId, statuses, int(req.Offset), 15)
	if err != nil {
		return nil, err
	}

	weight := map[v1.ProcessStatus]int{
		v1.ProcessStatus_StatusDone:    1,
		v1.ProcessStatus_StatusError:   99,
		v1.ProcessStatus_StatusRunning: 97,
		v1.ProcessStatus_StatusStop:    98,
		v1.ProcessStatus_StatusRetry:   3,
		v1.ProcessStatus_StatusReady:   4,
	}

	sort.Slice(processes, func(i, j int) bool {
		return weight[processes[i].Status]-weight[processes[j].Status] > 0
	})

	return &v1.ListProcessResponse{
		Processes: processes,
	}, nil
}
func (s *ProcessService) RetryProcess(ctx context.Context, req *v1.RetryProcessRequest) (*v1.RetryProcessResponse, error) {

	err := s.dispatcher.RetryProcessProfile(ctx, req.ProcessId, req.ProfileId, req.TaskId)
	if err != nil {
		return nil, err
	}

	return &v1.RetryProcessResponse{}, nil
}
func (s *ProcessService) GetProcessUpdatedAt(ctx context.Context, req *v1.GetProcessUpdatedAtRequest) (*v1.GetProcessUpdatedAtResponse, error) {

	at, err := s.processRepository.GetProcessUpdatedAt(ctx, req.GetProcessId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessUpdatedAtResponse{
		UpdatedAt: timestamppb.New(*at),
	}, nil
}
func (s *ProcessService) GetProcessTaskHistory(ctx context.Context, req *v1.GetProcessTaskHistoryRequest) (*v1.GetProcessTaskHistoryResponse, error) {
	records, err := s.processRepository.GetProcessTaskHistory(ctx, req.GetTaskId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessTaskHistoryResponse{
		Records: records,
	}, nil
}
func (s *ProcessService) StopProcess(ctx context.Context, req *v1.StopProcessRequest) (*v1.StopProcessResponse, error) {
	if err := s.dispatcher.StopProcess(ctx, req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.StopProcessResponse{}, nil
}
func (s *ProcessService) ResumeProcess(ctx context.Context, req *v1.ResumeProcessRequest) (*v1.ResumeProcessResponse, error) {
	//userId, err := user.GetUserId(ctx)
	//if err != nil {
	//	return nil, err
	//}
	s.dispatcher.ResumeProcess(ctx, req.ProcessId)
	return &v1.ResumeProcessResponse{}, nil
}
func (s *ProcessService) CancelProcess(ctx context.Context, req *v1.CancelProcessRequest) (*v1.CancelProcessResponse, error) {

	if err := s.dispatcher.KillProcess(ctx, req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.CancelProcessResponse{}, nil
}
func (s *ProcessService) EstimateCost(ctx context.Context, req *v1.EstimateCostRequest) (*v1.EstimateCostResponse, error) {
	res, err := s.dispatcher.EstimateTaskCost(ctx, req.ProfileId, req.TaskId)
	if err != nil {
		e := err.Error()
		return &v1.EstimateCostResponse{
			Error: &e,
		}, nil
	}
	return &v1.EstimateCostResponse{
		Error: nil,
		Data:  res,
	}, nil
}
func (s *ProcessService) GetTaskSettings(ctx context.Context, req *v1.GetTaskSettingsRequest) (*v1.GetTaskSettingsResponse, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	stgs, err := s.settingsService.GetSettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	gasLimit, ok := stgs.TaskGasLimitMap[req.TaskType.String()]
	if !ok {
		return &v1.GetTaskSettingsResponse{}, nil
	}

	return &v1.GetTaskSettingsResponse{
		GasLimit: defi.AmountUni(big.NewInt(gasLimit), req.Network),
	}, nil
}
func (s *ProcessService) SetTaskSettings(ctx context.Context, req *v1.SetTaskSettingsRequest) (*v1.SetTaskSettingsResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	stgs, err := s.settingsService.GetSettings(ctx, userId)
	if err != nil {
		return nil, err
	}

	if stgs.TaskGasLimitMap == nil {
		stgs.TaskGasLimitMap = map[string]int64{}
	}

	stgs.TaskGasLimitMap[req.TaskType.String()] = req.Wei

	if err := s.settingsService.UpdateSettings(ctx, stgs); err != nil {
		return nil, err
	}

	return &v1.SetTaskSettingsResponse{}, nil
}
func (s *ProcessService) GetTaskTransactions(ctx context.Context, req *v1.GetTaskTransactionsReq) (*v1.GetTaskTransactionsRes, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.processRepository.TransactionsByTaskId(ctx, userId, req.TaskId)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.Transaction, len(res))

	for i := range res {
		tmp[i] = res[i].ToPB()
	}

	return &v1.GetTaskTransactionsRes{
		Transactions: tmp,
	}, nil
}
func (s *ProcessService) GetProfileTransactions(ctx context.Context, req *v1.GetProfileTransactionsReq) (*v1.GetProfileTransactionsRes, error) {
	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.processRepository.TransactionsByProfileId(ctx, userId, req.ProfileId)
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.Transaction, len(res))

	for i := range res {
		tmp[i] = res[i].ToPB()
	}

	return &v1.GetProfileTransactionsRes{
		Transactions: tmp,
	}, nil
}
