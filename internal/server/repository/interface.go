package repository

import (
	"context"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error)
	ListProfiles(ctx context.Context, req *v1.ListProfileRequest) (*v1.ListProfileResponse, error)
	DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error)
	SearchProfile(ctx context.Context, req *v1.SearchProfileRequest) (*v1.SearchProfileResponse, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	ValidateLabel(ctx context.Context, request *v1.ValidateLabelRequest) (*bool, error)
	UpdateProfile(ctx context.Context, req *Profile) error
}

type WithdrawerRepository interface {
	CreateWithdrawer(ctx context.Context, req *v1.CreateWithdrawerRequest) (*v1.CreateWithdrawerResponse, error)
	ListWithdrawers(ctx context.Context, req *v1.ListWithdrawerRequest) (*v1.ListWithdrawerResponse, error)
	GetWithdrawers(ctx context.Context, id string) (*Withdrawer, error)
	DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error)
}

type FlowRepository interface {
	CreateFlow(ctx context.Context, req *v1.CreateFlowRequest) (*v1.CreateFlowResponse, error)
	ListFlows(ctx context.Context, req *v1.ListFlowRequest) (*v1.ListFlowResponse, error)
	DeleteFlow(ctx context.Context, req *v1.DeleteFlowRequest) (*v1.DeleteFlowResponse, error)
	GetFlow(ctx context.Context, req *v1.GetFlowRequest) (*v1.GetFlowResponse, error)
}

type ProcessRepository interface {
	CreateProcess(ctx context.Context, req *v1.CreateProcessRequest) (*v1.CreateProcessResponse, error)
	UpdateProcess(ctx context.Context, req *v1.Process) (*v1.Process, error)
	GetProcessArg(ctx context.Context, req *v1.GetProcessRequest) (*v1.GetProcessResponse, error)
	ListProcessByStatus(ctx context.Context, status v1.ProcessStatus) ([]*v1.Process, error)
	ListProcessByUser(ctx context.Context, userId string) ([]*v1.Process, error)
}

type SettingsRepository interface {
	GetSettings(ctx context.Context, userId string) (*v1.Settings, error)
	UpdateSettings(ctx context.Context, request *v1.Settings) error
}

type Repository interface {
	ProfileRepository
	WithdrawerRepository
	FlowRepository
	SettingsRepository
}
