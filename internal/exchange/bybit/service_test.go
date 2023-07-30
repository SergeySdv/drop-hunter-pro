package bybit

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {

	cfg := tests.GetConfig()
	c := NewService(&Config{
		ApiKey:    cfg.BybitKey,
		SecretKey: cfg.BybitSecret,
		HttpCli:   cfg.Cli,
	})

	b, err := c.GetBalance(context.Background(), "USDT")
	assert.NoError(t, err)
	assert.NotNil(t, b)
}
