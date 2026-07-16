package transaction

import (
	"banking_system/internal/domain/account"
	"banking_system/internal/domain/transaction"

	"github.com/google/uuid"
)

type TransactionService struct {
	accounts account.AccountRepsitory
	transactions transaction.TransactionRepository
}

func NewTransactionService(transactionRepo transaction.TransactionRepository, accts account.AccountRepsitory) *TransactionService {
	return &TransactionService{
		accounts: accts,
		transactions: transactionRepo,
	}
}

func (s *TransactionService) Transfer(
    fromAccountID uuid.UUID,
    toAccountID uuid.UUID,
    amount int,
    description string,
) error {
	sourceAcct, err := s.accounts.FindByID(fromAccountID)
	if err != nil {
		return  err
	}
	receivingAcct, err := s.accounts.FindByID(toAccountID)
	if err != nil {
		return err
	}

	// transaction
	sourceAcct.Balance -= amount
	receivingAcct.Balance += amount

	if err := s.accounts.Update(sourceAcct); err != nil {
		return err
	}
	if err := s.accounts.Update(receivingAcct); err != nil {
		return err
	}

	txn, err := transaction.New(fromAccountID, toAccountID, transaction.Transfer, amount, description)
	if err != nil {
		return err
	}

	return s.transactions.Save(*txn)
}