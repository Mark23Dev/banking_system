package accountstorage

import (
	"banking_system/internal/domain/account"
	"banking_system/internal/utils"
	"errors"

	"github.com/google/uuid"
)

type FileAccountsStore struct {
	filepath string
}

func NewFileAccountsStore(path string) *FileAccountsStore {
	return &FileAccountsStore{
		filepath: path,
	}
}

func (f *FileAccountsStore) Save(acct account.Account) error {
	accounts, err := f.FindAll()
	if err != nil {
		return err
	}

	accounts = append(accounts, acct)

	return utils.WriteJSON(f.filepath, accounts)
}

func (f *FileAccountsStore) FindAll() ([]account.Account, error) {
	var accounts []account.Account

	err := utils.ReadJSON(f.filepath, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (f *FileAccountsStore) FindByID(id uuid.UUID) (account.Account, error) {
	accounts, err := f.FindAll()
	if err != nil {
		return account.Account{}, err
	}

	for _, acct := range accounts {
		if acct.ID == id {
			return acct, nil
		}
	}

	return account.Account{}, errors.New("account not found")
}

func (f *FileAccountsStore) FindByAccountNumber(accountNumber string) (account.Account, error) {
	accounts, err := f.FindAll()
	if err != nil {
		return account.Account{}, err
	}

	for _, acct := range accounts {
		if acct.AccountNumber == accountNumber {
			return acct, nil
		}
	}

	return account.Account{}, errors.New("account not found")
}

func (f *FileAccountsStore) Update(updated account.Account) error {
	accounts, err := f.FindAll()
	if err != nil {
		return err
	}

	for i, acct := range accounts {
		if acct.ID == updated.ID {
			accounts[i] = updated
			return utils.WriteJSON(f.filepath, accounts)
		}
	}

	return errors.New("account not found")
}