package snapshot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	s := NewService(&Config{
		Host: "http://localhost:3344",
	})

	_, err := s.ActiveProposals(context.Background(), &ActiveProposalsReq{
		ProviderRPC: "https://arb1.arbitrum.io/rpc",
		Space:       "stgdao.eth",
		Pk:          "",
	})
	assert.NoError(t, err)
}
