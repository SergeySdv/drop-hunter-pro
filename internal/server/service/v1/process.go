package v1

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/process"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/server/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProcessService struct {
	v1.UnimplementedProcessServiceServer
	repository *repository.PGRepository
	dispatcher *process.Dispatcher
}

func NewProcessService(repository *repository.PGRepository, dispatcher *process.Dispatcher) *ProcessService {
	return &ProcessService{
		repository: repository,
		dispatcher: dispatcher,
	}
}

func (s *ProcessService) CreateProcess(ctx context.Context, req *v1.CreateProcessRequest) (*v1.CreateProcessResponse, error) {
	res, err := s.repository.CreateProcess(ctx, req)
	if err != nil {
		return nil, err
	}
	go s.dispatcher.StartProcess(res.Process.Id)
	return res, nil
}
func (s *ProcessService) GetProcess(ctx context.Context, req *v1.GetProcessRequest) (*v1.GetProcessResponse, error) {
	return s.repository.GetProcessArg(ctx, req)
}
func (s *ProcessService) ListProcess(ctx context.Context, req *v1.ListProcessRequest) (*v1.ListProcessResponse, error) {
	processes, err := s.repository.ListProcessByUser(ctx, user.GetUserId(ctx))
	if err != nil {
		return nil, err
	}

	return &v1.ListProcessResponse{
		Processes: processes,
	}, nil
}

func (s *ProcessService) RetryProcess(ctx context.Context, req *v1.RetryProcessRequest) (*v1.RetryProcessResponse, error) {

	err := s.dispatcher.RetryProcessProfile(ctx, req.ProcessId, req.TaskId)
	if err != nil {
		return nil, err
	}

	return &v1.RetryProcessResponse{}, nil
}

func (s *ProcessService) GetProcessUpdatedAt(ctx context.Context, req *v1.GetProcessUpdatedAtRequest) (*v1.GetProcessUpdatedAtResponse, error) {

	at, err := s.repository.GetProcessUpdatedAt(ctx, req.GetProcessId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessUpdatedAtResponse{
		UpdatedAt: timestamppb.New(*at),
	}, nil
}

func (s *ProcessService) GetProcessTaskHistory(ctx context.Context, req *v1.GetProcessTaskHistoryRequest) (*v1.GetProcessTaskHistoryResponse, error) {
	records, err := s.repository.GetProcessTaskHistory(ctx, req.GetTaskId())
	if err != nil {
		return nil, err
	}

	return &v1.GetProcessTaskHistoryResponse{
		Records: records,
	}, nil
}

func (s *ProcessService) StopProcess(ctx context.Context, req *v1.StopProcessRequest) (*v1.StopProcessResponse, error) {
	if err := s.dispatcher.StopProcess(req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.StopProcessResponse{}, nil
}

func (s *ProcessService) ResumeProcess(ctx context.Context, req *v1.ResumeProcessRequest) (*v1.ResumeProcessResponse, error) {
	s.dispatcher.ResumeProcess(req.ProcessId)
	return &v1.ResumeProcessResponse{}, nil
}

func (s *ProcessService) CancelProcess(ctx context.Context, req *v1.CancelProcessRequest) (*v1.CancelProcessResponse, error) {

	if err := s.dispatcher.KillProcess(req.ProcessId); err != nil {
		return nil, err
	}

	return &v1.CancelProcessResponse{}, nil
}
