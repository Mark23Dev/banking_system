package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID
	FromAccountID uuid.UUID
	ToAccountID uuid.UUID
	RefNumber string
	Type TransactionType
	Status
	Amount int
	Description string
	CreatedAt time.Time
}

type TransactionType int

const (
    Deposit TransactionType = iota
    Withdrawal
    Transfer
)

type Status int

const (
    Pending Status = iota
    Completed
    Failed
    Reversed
)

func New(fromAccountID, toAccountID uuid.UUID, transactionType TransactionType, amount int, description string) (*Transaction, error) {
	refNumber := uuid.NewString()
	
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	now := time.Now()

	return &Transaction{
		ID: uuid.New(),
		FromAccountID: fromAccountID,
		ToAccountID: toAccountID,
		RefNumber: refNumber,
		Type: transactionType,
		Status: Pending,
		Amount: amount,
		Description: description,
		CreatedAt: now,
	}, nil
}

func (t *Transaction) Complete() {
	t.Status = Completed
}

func (t *Transaction) Fail() {
	t.Status = Failed
}

func (t *Transaction) Reverse() error {
	if t.Status != Completed {
		return ErrTransactionNotCompleted
	}

	t.Status = Reversed
	return nil
}