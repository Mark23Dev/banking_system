package account

import "errors"

var (
	ErrInadequateBalance =    errors.New("Inadequate balance")
	ErrAccountFrozen     =    errors.New("Account frozen")
	ErrAccountClosed     =    errors.New("Account closed")
	ErrInvalidAmount     =    errors.New("invalid amount")
	ErrAccountHasBalance =    errors.New("Account has balance")
	ErrAccountInactive   =    errors.New("account not active")
	ErrAccountAlreadyActive = errors.New("account already active")
)