package process

import (
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"sort"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

type Process struct {
	p  *v1.Process
	mu *sync.Mutex
}

func NewProcess(p *v1.Process) *Process {
	return &Process{
		mu: &sync.Mutex{},
		p:  p,
	}
}

func (p *Process) Copy() (*Process, error) {

	marshal, err := protojson.Marshal(p.p)
	if err != nil {
		return nil, err
	}

	var t v1.Process
	if err := protojson.Unmarshal(marshal, &t); err != nil {
		return nil, err
	}

	return NewProcess(&t), nil
}

func (p *Process) NeedToExec() bool {
	return p.p.Status != v1.ProcessStatus_StatusDone
}

type Profile struct {
	p *v1.ProcessProfile
}

func (p *Profile) Copy() (*Profile, error) {
	marshal, err := protojson.Marshal(p.p)
	if err != nil {
		return nil, err
	}

	var t v1.ProcessProfile
	if err := protojson.Unmarshal(marshal, &t); err != nil {
		return nil, err
	}

	return NewProfile(&t), nil
}

func NewProfile(p *v1.ProcessProfile) *Profile {
	sort.Slice(p.Tasks, func(i, j int) bool {
		return p.Tasks[i].Task.Weight < p.Tasks[j].Task.Weight
	})

	return &Profile{p: p}
}

func (p *Profile) Status() v1.ProcessStatus {
	return p.p.Status
}

func (p *Profile) NeedToExec() bool {
	return p.p.Status != v1.ProcessStatus_StatusDone
}

var ErrNoTaskToExec = errors.New("no task to exec")

func (p *Profile) GetTaskToExec() (*v1.ProcessTask, error) {

	for _, task := range p.p.Tasks {

		if task.Status == v1.ProcessStatus_StatusRunning {
			return task, nil
		}

		if task.Status == v1.ProcessStatus_StatusError {
			return nil, ErrNoTaskToExec
		}

		if task.Status == v1.ProcessStatus_StatusReady {
			return task, nil
		}

		if task.Status == v1.ProcessStatus_StatusRetry {
			return task, nil
		}
	}

	return nil, ErrNoTaskToExec
}
