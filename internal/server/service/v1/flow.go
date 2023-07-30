package v1

import (
	"context"

	"github.com/google/uuid"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
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

func (s *FlowService) UpdateFlow(ctx context.Context, req *v1.UpdateFlowRequest) (*v1.UpdateFlowResponse, error) {

	userId, err := user.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	parentFlowId := req.Flow.Id

	req.Flow.Id = uuid.New().String()
	var f repository.Flow
	if err := f.FromPB(req.Flow, userId); err != nil {
		return nil, err
	}

	if err := s.repository.UpdateFlow(ctx, parentFlowId, &f); err != nil {
		return nil, err
	}

	flow, err := s.repository.GetFlow(ctx, &v1.GetFlowRequest{Id: req.Flow.Id})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateFlowResponse{
		Flow: flow.Flow,
	}, nil
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
