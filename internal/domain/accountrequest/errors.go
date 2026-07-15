package accountrequest

import "errors"

var (
	ErrRequestNotPending = errors.New("request already executed")
)