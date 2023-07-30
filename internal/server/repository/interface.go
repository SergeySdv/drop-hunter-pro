package repository

import (
	"context"
	"time"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

// todo: исключить из зависимостей v1

type ProfileRepository interface {
	CreateProfile(ctx context.Context, req *Profile) error
	ListProfiles(ctx context.Context, userId string) ([]Profile, error)
	DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error)
	SearchProfile(ctx context.Context, pattern, userId string) ([]Profile, error)
	SearchNotConnectedOkexDepositProfile(ctx context.Context, userId string) ([]Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	ValidateLabel(ctx context.Context, request *ValidateLabelReq) (*bool, error)
	UpdateProfile(ctx context.Context, req *Profile) error
	ExportProfiles(ctx context.Context, userId string) ([]Profile, error)
}

type WithdrawerRepository interface {
	CreateWithdrawer(ctx context.Context, req *Withdrawer) (*Withdrawer, error)
	CreateSubWithdrawer(ctx context.Context, req *Withdrawer) error
	ListWithdrawers(ctx context.Context, userId string) ([]Withdrawer, error)
	ListSubWithdrawers(ctx context.Context, id, userId string) ([]Withdrawer, error)
	DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error)
	GetWithdrawer(ctx context.Context, id, userId string) (*Withdrawer, error)
	UpdateWithdrawer(ctx context.Context, req *Withdrawer) error
	ExportExchangeAccounts(ctx context.Context, userId string) ([]Withdrawer, error)

	OkexDepositAddrAttach(ctx context.Context, req *OkexDepositAddr) error
	OkexDepositAddrDetach(ctx context.Context, req *OkexDepositAddr) error
	ListDepositAddrAttach(ctx context.Context, userId string) (map[Addr]OkexDepositAddr, error)
	GetAttachedAddr(ctx context.Context, profileId, userId string) (*OkexDepositAddr, error)
}

type StatRepository interface {
	GetDailyUserImpact(ctx context.Context, userId string) (*int64, error)
	GetDailyTopImpact(ctx context.Context) (*int64, error)
	GetDailyTotalImpact(ctx context.Context) (*int64, error)
}

type FlowRepository interface {
	CreateFlow(ctx context.Context, req *v1.CreateFlowRequest) (*v1.CreateFlowResponse, error)
	ListFlows(ctx context.Context, req *v1.ListFlowRequest) (*v1.ListFlowResponse, error)
	DeleteFlow(ctx context.Context, req *v1.DeleteFlowRequest) (*v1.DeleteFlowResponse, error)
	GetFlow(ctx context.Context, req *v1.GetFlowRequest) (*v1.GetFlowResponse, error)
	UpdateFlow(ctx context.Context, parentFlowId string, req *Flow) (err error)
}

type ProcessRepository interface {
	CreateProcess(ctx context.Context, req *ProcessArg) error
	GetProcessArg(ctx context.Context, req *v1.GetProcessRequest, userId string) (*v1.GetProcessResponse, error)
	ListProcessByUser(ctx context.Context, userId string, statuses []string, offset, limit int) ([]*v1.Process, error)
	GetProcessUpdatedAt(ctx context.Context, processId string) (*time.Time, error)
	UpdateProcessStatus(ctx context.Context, req *UpdateProcess) error
	GetProcessStatus(ctx context.Context, processId string) (*v1.ProcessStatus, error)
	DeleteProcess(ctx context.Context, id string) error
	ListProcessIdsByStatus(ctx context.Context, statuses ...v1.ProcessStatus) ([]string, error)
	ListProcessIdsForAutoRetry(ctx context.Context) ([]string, error)
	GetProcessUser(ctx context.Context, processId string) (*string, error)
	UpdateProcessTime(ctx context.Context, id string) error
	UpdateProcessAutoRetry(ctx context.Context, id string, enable bool) error

	ProcessProfileRepository
	ProcessTaskRepository
	TransactionRepository
}

type TransactionRepository interface {
	TransactionAdd(ctx context.Context, req *Transaction) error
	TransactionsByTaskId(ctx context.Context, userId, taskId string) ([]Transaction, error)
	TransactionsByProfileId(ctx context.Context, userId, profileId string) ([]Transaction, error)
}

type ProcessTaskRepository interface {
	GetProcessTask(ctx context.Context, taskId string) (*ProcessTask, error)
	GetProcessTaskHistory(ctx context.Context, taskId string) ([]*v1.ProcessTaskHistoryRecord, error)
	GetTaskStatus(ctx context.Context, taskId string) (*v1.ProcessStatus, error)
	UpdateProcessTaskStatus(ctx context.Context, status, taskId, processId string) error
	UpdateProcessTask(ctx context.Context, req *v1.ProcessTask, taskId, processId, profileId string) error
	RecordStatusChanged(ctx context.Context, taskId string, before, after v1.ProcessStatus, msg ...string) error
}

type ProcessProfileRepository interface {
	GetProcessProfileArg(ctx context.Context, profileId string) (*ProcessProfileArg, error)
	UpdateProcessProfileStatus(ctx context.Context, profileId, processId string, status string) error
	GetProcessProfileTaskStatuses(ctx context.Context, profileId string) ([]v1.ProcessStatus, error)
	GetProcessProfileStatuses(ctx context.Context, processId string) ([]v1.ProcessStatus, error)
	GetProcessProfileIds(ctx context.Context, processId string) ([]string, error)
	GetProcessProfileStatus(ctx context.Context, profileId string) (*v1.ProcessStatus, error)
}

type SettingsRepository interface {
	GetSettings(ctx context.Context, userId string) (*v1.Settings, error)
	UpdateSettings(ctx context.Context, request *v1.Settings) error
}

type UserRepository interface {
	GetOrCreateUser(ctx context.Context, user *User) (*User, bool, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	SubscribeAlerts(ctx context.Context, email, chatId string) error
	GetUserTelegramChatId(ctx context.Context, userId string) (*string, error)
	// todo: unsubscribe alerts
}

type Repository interface {
	ProfileRepository
	WithdrawerRepository
	FlowRepository
	SettingsRepository
	UserRepository
	ProcessRepository
	StatRepository
}
