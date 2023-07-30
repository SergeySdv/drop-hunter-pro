package scan

import (
	"context"
	"time"
)

type Scanner interface {
	WaitComplete(ctx context.Context, txId string, interval time.Duration) (*TxStatus, error)
}
