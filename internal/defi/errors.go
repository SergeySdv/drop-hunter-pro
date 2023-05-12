package defi

import (
	"github.com/pkg/errors"
)

func errTokenNotSupported(t Token) error {
	return errors.Wrap(ErrTokenNotSupported, "token: "+string(t))
}

var (
	ErrTokenNotSupported = errors.New("token is not supported in this network")
)
