package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	PINHash      string

	Role
	Status

	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create a new user
func New(username, email string, passwordHash, pinHash string) *User {
	now := time.Now()

	return &User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: passwordHash,
		PINHash:      pinHash,
		Role:         Guest,
		Status:       Pending,
		Email:	      email,	
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// Approve a pending user
func (u *User) Approve() error {
	if u.Status != Pending {
		return ErrUserNotPending
	}

	u.Status = Accepted
	u.touch()

	return nil
}

// Reject a pending user
func (u *User) Reject() error {
	if u.Status != Pending {
		return ErrUserNotPending
	}

	u.Status = Rejected
	u.touch()

	return nil
}

// convert an approved guest into a customer
func (u *User) PromoteToCustomer() error {
	if u.Status != Accepted {
		return ErrUserNotApproved
	}

	if u.Role != Guest {
		return ErrInvalidRoleTransition
	}

	u.Role = Customer
	u.touch()

	return nil
}

// Propmote to manager
// func (u *User) PromoteToManager() error {
// 	if u.Role != 
// }

// change a user's role
func (u *User) ChangeRole(role Role) {
	u.Role = role
	u.touch()
}

// update the password hash
func (u *User) ChangePassword(hash string) {
	u.PasswordHash = hash
	u.touch()
}

// update the PIN hash
func (u *User) ChangePIN(hash string) {
	u.PINHash = hash
	u.touch()
}

func (u User) IsAdmin() bool {
	return u.Role == Admin
}

func (u User) IsManager() bool {
	return u.Role == Manager
}

func (u User) IsCustomer() bool {
	return u.Role == Customer
}

func (u User) IsGuest() bool {
	return u.Role == Guest
}

func (u User) IsPending() bool {
	return u.Status == Pending
}

func (u User) IsAccepted() bool {
	return u.Status == Accepted
}

func (u User) IsRejected() bool {
	return u.Status == Rejected
}

func (u *User) touch() {
	u.UpdatedAt = time.Now()
}