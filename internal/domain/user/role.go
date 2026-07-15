package user


type Role int

const (
	Admin Role = iota
	Guest
	Customer
	Manager

)