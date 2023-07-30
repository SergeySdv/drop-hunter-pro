package process

import (
	"sync"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
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
