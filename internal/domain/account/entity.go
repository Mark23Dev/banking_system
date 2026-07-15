package account

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
	AccountType
	AccountNumber string
	Status
	Balance int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Status int

const (
	Active Status = iota
	Frozen
	Closed
)

type AccountType int

const (
	Checking AccountType = iota
	Savings
	Specialty
	Business
)

func GenerateAccountNumber(digits int) (string, error) {
	// Set the bounds for random number generation (e.g., for 10 digits, max is 9999999999)
	maxLimit := new(big.Int)
	maxLimit.Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)

	// Generate secure random number
	val, err := rand.Int(rand.Reader, maxLimit)
	if err != nil {
		return "", err
	}

	// Format with leading zeros if it's shorter than the desired length
	formatStr := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(formatStr, val), nil
}


func New(customerID uuid.UUID, accountType AccountType) (*Account , error) {
	now := time.Now()
	accountNumber, err := GenerateAccountNumber(10)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:            uuid.New(),
		CustomerID:    customerID,
		AccountType:   accountType,
		AccountNumber: accountNumber,
		Balance:       0,
		Status:        Active,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

func (a *Account) Deposit(amount int) error {
	if a.Status == Frozen {
		return ErrAccountFrozen
	}
	if !a.IsActive() {
		return ErrAccountClosed
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount
	a.touch()
	return nil
}
func (a *Account) Withdraw(amount int) error {
	if a.Status == Frozen {
		return ErrAccountFrozen
	}
	if !a.IsActive() {
		return ErrAccountClosed
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if !a.CanWithdraw(amount) {
		return ErrInadequateBalance
	}

	a.Balance -= amount
	a.touch()
	return nil
}

func (a *Account) Freeze() {
	a.Status = Frozen
	a.touch()
}

func (a *Account) Activate() {
	a.Status = Active
	a.touch()
}

func (a *Account) CanWithdraw(amount int) bool {
	return a.Balance >= amount
}

func (a *Account) Close() error {
	if a.Balance > 0 {
		return ErrAccountHasBalance
	}
	a.Status = Closed
	a.touch()
	return nil
}

func (a *Account) IsActive() bool {
	return a.Status == Active
}

func (a *Account) IsFrozen() bool {
	return a.Status == Frozen
}

func (a *Account) touch() {
	a.UpdatedAt = time.Now()
}