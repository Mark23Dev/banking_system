package account

import "github.com/google/uuid"

type AccountRepsitory interface {
	Save(account Account) error
	FindAll() ([]Account, error)
	FindByID(id uuid.UUID) (Account, error)
	FindByAccountNumber(accountNumber string) (Account, error)
	Update(updated Account) error
}