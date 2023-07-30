package task

import (
	"fmt"

	"github.com/pkg/errors"
)

func ErrGasIsOverMax(max, estimated string) error {
	return errors.New(fmt.Sprintf("gas (%s USD) is higher than max (%s USD)", estimated, max))
}

var (
	ErrUserHasNoBalance              = errors.New("User has not enough balance. visit billing page please")
	ErrProfileHasNoConnectedOkexAddr = errors.New("Profile is not connected to okex deposit wallet")
)
