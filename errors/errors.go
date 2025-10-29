package errors

import (
	errs "errors"
	"fmt"
)

var ErrTelebot = errs.New("telebot error")

var (
	ErrBadRecipient         = fmt.Errorf("%w: recipient is invalid", ErrTelebot)
	ErrUnsupportedWhat      = fmt.Errorf("%w: unsupported what argument", ErrTelebot)
	ErrNothingToEdit        = fmt.Errorf("%w: nothing to edit", ErrTelebot)
	ErrContextInsufficient  = fmt.Errorf("%w: context is insufficient", ErrTelebot)
	ErrInvalidParam         = fmt.Errorf("%w: invalid parameter", ErrTelebot)
	ErrTelegramInternal     = fmt.Errorf("%w: telegram internal error", ErrTelebot)
	ErrFlood                = fmt.Errorf("%w: flood error", ErrTelebot)
	ErrRequestFailed        = fmt.Errorf("%w: request failed", ErrTelebot)
	ErrRequestResultIsEmpty = fmt.Errorf("%w: request result is empty", ErrTelebot)
)

// wrapError returns new wrapped telebot-related error.
func Wrap(err error) error {
	return fmt.Errorf("%w: %w", ErrTelebot, err)
}
