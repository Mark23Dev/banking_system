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

func (a *AccountService) CreateNewAccount(customerID uuid.UUID, accountType account.AccountType) error {
	acct, err := account.New(customerID, accountType)
	if err != nil {
		return err
	}
	return a.repo.Save(*acct)
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

func (a *AccountService) DepositToAccount(accountID uuid.UUID, amount int) error {
	acct, err := a.repo.FindByID(accountID)
	if err != nil {
		return err
	}
	if err := acct.Deposit(amount); err != nil {
		return err
	}

	return a.repo.Update(acct)
}

func (a *AccountService) WithdrawFromAccount(accountID uuid.UUID, amount int) error {
	acct, err := a.repo.FindByID(accountID)
	if err != nil {
		return err
	}
	if err := acct.Withdraw(amount); err != nil {
		return err
	}
	return a.repo.Update(acct)
}


func (a *AccountService) FreezeAccount(accountID uuid.UUID) error {
	acct, err := a.repo.FindByID(accountID)
	if err != nil {
		return err
	}
	if !acct.IsActive() {
		return account.ErrAccountInactive
	}
	acct.Freeze()
	return a.repo.Update(acct)

}

func (a *AccountService) ActivateAccount(accountID uuid.UUID) error {
	acct, err := a.repo.FindByID(accountID)
	if err != nil {
		return err
	}
	if acct.IsActive() {
		return account.ErrAccountAlreadyActive
	}
	acct.Activate()
	return a.repo.Update(acct)

}

func (a *AccountService) CloseAccount(accountID uuid.UUID) error {
	acct, err := a.repo.FindByID(accountID)
	if err != nil {
		return err
	}
	if !acct.IsActive() {
		return account.ErrAccountInactive
	}
	acct.Close()
	return a.repo.Update(acct)
}

