package zksynclite

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
)

func (c *Client) WaitTxComplete(ctx context.Context, tx common.Hash) error {
	ticker := time.NewTicker(time.Second * 5)

	for {
		tx, err := c.TxInfo(ctx, tx)
		if err != nil {
			return err
		}

		if !tx.Success {
			return defi.ErrTxStatusFailed
		}

		if tx.FailReason != nil {
			return defi.ErrTxStatusFailed
		}

		if tx.Block.Verified {
			return nil
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *Client) TxInfo(ctx context.Context, hash common.Hash) (*TxInfoRes, error) {
	rpccli := c.ethcli.Client()
	var a TxInfoRes
	err := rpccli.CallContext(ctx, &a, "tx_info", hash)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

type TxInfoRes struct {
	Executed   bool        `json:"executed"`
	Success    bool        `json:"success"`
	FailReason interface{} `json:"failReason"`
	Block      struct {
		BlockNumber int  `json:"blockNumber"`
		Committed   bool `json:"committed"`
		Verified    bool `json:"verified"`
	} `json:"block"`
}
