package accountrequest

import (
	"banking_system/internal/domain/account"
	"time"

	"github.com/google/uuid"
)

type AccountRequest struct {
	ID uuid.UUID

	UserID uuid.UUID

	AccountType account.AccountType
	AccountID *uuid.UUID

	Status

	ReviewedBy *uuid.UUID
	ReviewedAt *time.Time

	CreatedAt time.Time
}


type Status int


const (
	Pending Status = iota
	Approved
	Rejected
)

func New(userID uuid.UUID, accountType account.AccountType) *AccountRequest {
	now := time.Now()

	return &AccountRequest{
		ID: uuid.New(),
		UserID: userID,
		AccountType: accountType,
		AccountID: nil,
		Status: Pending,
		ReviewedBy: nil,
		ReviewedAt: nil,
		CreatedAt: now,
	}
}

func (a *AccountRequest) Approve(managerID uuid.UUID) error {
	if a.Status != Pending {
		return ErrRequestNotPending
	}

	now := time.Now()

	a.Status = Approved
	a.ReviewedBy = &managerID
	a.ReviewedAt = &now

	return nil
}

func (a *AccountRequest) Reject(managerID uuid.UUID) error {
	if a.Status != Pending {
		return ErrRequestNotPending
	}
	now := time.Now()

	a.Status = Rejected
	a.ReviewedAt = &now
	a.ReviewedBy = &managerID
	return nil

}