package transactionstorage

import (
	"banking_system/internal/domain/transaction"
	"banking_system/internal/utils"
	"errors"

	"github.com/google/uuid"
)

type FileTransactionsStore struct {
	filepath string
}

func NewFileTransactionsStore(path string) *FileTransactionsStore {
	return &FileTransactionsStore{
		filepath: path,
	}
}

func (f *FileTransactionsStore) Save(tx transaction.Transaction) error {
	transactions, err := f.FindAll()
	if err != nil {
		return err
	}

	transactions = append(transactions, tx)

	return utils.WriteJSON(f.filepath, transactions)
}

func (f *FileTransactionsStore) FindByID(id uuid.UUID) (transaction.Transaction, error) {
	transactions, err := f.FindAll()
	if err != nil {
		return transaction.Transaction{}, err
	}

	for _, tx := range transactions {
		if tx.ID == id {
			return tx, nil
		}
	}

	return transaction.Transaction{}, errors.New("transaction not found")
}

func (f *FileTransactionsStore) FindAll() ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction

	err := utils.ReadJSON(f.filepath, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (f *FileTransactionsStore) TransactionsByAccount(accountID uuid.UUID) ([]transaction.Transaction, error) {
	transactions, err := f.FindAll()
	if err != nil {
		return nil, err
	}

	var result []transaction.Transaction

	for _, tx := range transactions {
		if tx.FromAccountID == accountID || tx.ToAccountID == accountID {
			result = append(result, tx)
		}
	}

	return result, nil
}