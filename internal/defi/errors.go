package defi

import (
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func NewErrTokenNotSupported(t v1.Token) error {
	return errors.Wrap(ErrTokenNotSupported, "token: "+string(t))
}

var (
	ErrTokenNotSupported = errors.New("token is not supported in this network")
)

var ErrTxStatusFailed = errors.New("transaction failed")
var ErrTxNotFound = errors.New("transaction is not found")
