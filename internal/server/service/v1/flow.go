package v1

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/repository"
)

type FlowService struct {
	v1.UnimplementedFlowServiceServer
	repository repository.FlowRepository
}

func NewFlowService(repository repository.FlowRepository) *FlowService {
	return &FlowService{
		repository: repository,
	}
}

func (s *FlowService) CreateFlow(ctx context.Context, req *v1.CreateFlowRequest) (*v1.CreateFlowResponse, error) {
	return s.repository.CreateFlow(ctx, req)
}
func (s *FlowService) ListFlow(ctx context.Context, req *v1.ListFlowRequest) (*v1.ListFlowResponse, error) {
	return s.repository.ListFlows(ctx, req)
}
func (s *FlowService) DeleteFlow(ctx context.Context, req *v1.DeleteFlowRequest) (*v1.DeleteFlowResponse, error) {
	s.repository.DeleteFlow(ctx, req)
	return &v1.DeleteFlowResponse{}, nil
}

func (s *FlowService) GetFlow(ctx context.Context, req *v1.GetFlowRequest) (*v1.GetFlowResponse, error) {
	return s.repository.GetFlow(ctx, req)
}
