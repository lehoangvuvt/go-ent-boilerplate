package transactiondomain

import "errors"

var (
	ErrInvalidTransition = errors.New("invalid transaction state transition")
	ErrMissingDetails    = errors.New("missing payment method details")
)
