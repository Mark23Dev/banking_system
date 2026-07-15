package auth

import "errors"

var (
	ErrInvalidUserCredentials = errors.New("invalid user credentials")
	ErrEmailAlreadyUsed     = errors.New("email already used")
	ErrUsernameAlreadyUsed  = errors.New("username already used")
)