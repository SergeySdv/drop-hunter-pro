package defi

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (c *EtheriumClient) WaitTxComplete(ctx context.Context, tx common.Hash) error {
	for {
		_, isPending, err := c.cli.TransactionByHash(ctx, tx)
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}
		if isPending {
			time.Sleep(time.Second * 5)
			continue
		}

		break
	}

	return nil
}
