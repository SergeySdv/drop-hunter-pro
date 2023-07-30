package lzscan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	s := NewService()
	_, err := s.WaitConfirm(context.Background(), "0x20c94fd02a32385e9233343bdef42dd5e1e438a6557b00dd136560ef354b918ad")
	assert.NoError(t, err)
}
