package task

import (
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/config"
)

func IsPayableTask(t v1.TaskType) bool {

	if config.CFG.Env == config.Local {
		return true
	}

	for _, tt := range PayableTasks {
		if t == tt {
			return true
		}
	}
	return false
}

func NeedPay(before, after *v1.ProcessTask) bool {
	if before.Status == after.Status {
		return false
	}

	if after.Status == v1.ProcessStatus_StatusDone {
		return false
	}

	if !IsPayableTask(after.Task.TaskType) {
		return false
	}

	return true
}
