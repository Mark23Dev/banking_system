package user

import "errors"

var (
	ErrUserNotPending      = errors.New("user is not pending")
	ErrUserNotApproved     = errors.New("user is not approved")
	ErrInvalidRoleTransition = errors.New("invalid role transition")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidPIN          = errors.New("invalid pin")
	ErrInvalidPassword     = errors.New("invalid password")
)