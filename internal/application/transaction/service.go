package transaction

import (
	"banking_system/internal/domain/account"
	"banking_system/internal/domain/transaction"

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
    fromAccountNumber string,
    toAccountNumber string,
    amount int,
    description string,
) error {
	sourceAcct, err := s.accounts.FindByAccountNumber(fromAccountNumber)
	if err != nil {
		return  err
	}
	receivingAcct, err := s.accounts.FindByAccountNumber(toAccountNumber)
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

	txn, err := transaction.New(fromAccountNumber, toAccountNumber, transaction.Transfer, amount, description)
	if err != nil {
		return err
	}

	return s.transactions.Save(*txn)
}