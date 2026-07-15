package user

type Status int

const (
	Accepted Status = iota
	Rejected
	Pending
)