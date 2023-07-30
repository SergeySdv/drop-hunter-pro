package defi

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (c *EtheriumClient) WaitTxComplete(ctx context.Context, txId common.Hash) error {

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		tx, isPending, err := c.Cli.TransactionByHash(ctx, txId)
		if err == nil {
			if !isPending {
				rec, err := bind.WaitMined(context.Background(), c.Cli, tx)
				if err != nil {
					return err
				}

				if rec.Status == types.ReceiptStatusFailed {
					return errors.Wrap(ErrTxStatusFailed, "invalid status")
				}
			}
		}

		if isPending {
			continue
		}

		if err != nil {
			if err.Error() == "not found" {
				return errors.Wrap(ErrTxNotFound, "blockchain error")
			}
		}

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:

			return nil
		}
	}
}
