package account

import (
	"banking_system/internal/domain/account"

	"github.com/google/uuid"
)

type AccountService struct {
	repo account.AccountRepsitory
}

func NewAccountService(r account.AccountRepsitory) *AccountService {
	return &AccountService{repo: r}
}

func (a *AccountService) CreateNewAccount(customerID uuid.UUID, accountType account.AccountType) (*account.Account, error) {
	acct, err := account.New(customerID, accountType)
	if err != nil {
		return nil, err
	}
	return acct, a.repo.Save(*acct)
}

func (a *AccountService) AccountsByCustomer(customerID uuid.UUID) ([]account.Account, error) {
	var customerAccts []account.Account
	accts, err := a.repo.FindAll()
	if err != nil {
		return nil, err
	}
	for _,acct := range accts {
		if acct.CustomerID == customerID {
			customerAccts = append(customerAccts, acct)
		}
	}
	return customerAccts, nil
}

func (a *AccountService) DepositToAccount(accountNumber string, amount int) error {
	acct, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return err
	}
	if err := acct.Deposit(amount); err != nil {
		return err
	}

	return a.repo.Update(acct)
}

func (a *AccountService) WithdrawFromAccount(accountNumber string, amount int) error {
	acct, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return err
	}
	if err := acct.Withdraw(amount); err != nil {
		return err
	}
	return a.repo.Update(acct)
}


func (a *AccountService) FreezeAccount(accountNumber string) error {
	acct, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return err
	}
	if !acct.IsActive() {
		return account.ErrAccountInactive
	}
	acct.Freeze()
	return a.repo.Update(acct)

}

func (a *AccountService) ActivateAccount(accountNumber string) error {
	acct, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return err
	}
	if acct.IsActive() {
		return account.ErrAccountAlreadyActive
	}
	acct.Activate()
	return a.repo.Update(acct)

}

func (a *AccountService) CloseAccount(accountNumber string) error {
	acct, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return err
	}
	if !acct.IsActive() {
		return account.ErrAccountInactive
	}
	acct.Close()
	return a.repo.Update(acct)
}

func (a *AccountService) Statement(accountNumber string) (*account.Account, error) {
	acc, err := a.repo.FindByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}