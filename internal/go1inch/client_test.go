package go1inch

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	c := NewClient(&Config{
		//HttpCli: tests.GetConfig().Cli,
	})

	res, _, err := c.ApproveSpender(context.Background(), Arbitrum)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
