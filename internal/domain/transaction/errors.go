package transaction

import "errors"

var (
	ErrInvalidAmount           = errors.New("invalid amount")
	ErrTransactionNotCompleted = errors.New("transaction not completed")
)