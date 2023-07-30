package process

import (
	"sort"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/pkg/errors"
)

type PP struct {
	pp   *v1.ProcessProfile
	ppdb *repository.ProcessProfileArg
}

var ErrNoTaskToExec = errors.New("no task to exec")

func NewPP(pp *repository.ProcessProfileArg) (*PP, error) {

	p, err := pp.ToPB()
	if err != nil {
		return nil, err
	}

	sort.Slice(p.Tasks, func(i, j int) bool {
		return p.Tasks[i].Task.Weight < p.Tasks[j].Task.Weight
	})

	return &PP{pp: p, ppdb: pp}, nil
}

func (p *PP) ToDB() *v1.ProcessProfile {
	return p.pp
}

func (p *PP) SetStatus(s v1.ProcessStatus) {
	p.pp.Status = s
}

func (p *PP) ProcessId() processId {
	return p.ppdb.ProcessId
}
func (p *PP) TaskStatuses() []v1.ProcessStatus {

	out := make([]v1.ProcessStatus, 0)

	if p == nil {
		return out
	}

	for _, task := range p.pp.Tasks {
		out = append(out, task.Status)
	}
	return out
}

func (p *PP) Status() v1.ProcessStatus {
	return p.pp.Status
}

func (p *PP) NeedToExec() bool {
	return p.pp.Status != v1.ProcessStatus_StatusDone
}

func (p *PP) GetTaskToExec() (*v1.ProcessTask, error) {

	for _, task := range p.pp.Tasks {

		switch task.Status {
		case v1.ProcessStatus_StatusRunning,
			v1.ProcessStatus_StatusReady,
			v1.ProcessStatus_StatusRetry:
			return task, nil
		case v1.ProcessStatus_StatusError, v1.ProcessStatus_StatusStop:
			return nil, ErrNoTaskToExec
		}
	}

	return nil, ErrNoTaskToExec
}

func (p *PP) ProfileFinished() bool {
	count := 0
	for _, task := range p.pp.Tasks {
		if task.Status == v1.ProcessStatus_StatusDone {
			count++
		}
	}

	return count == len(p.pp.Tasks)
}
