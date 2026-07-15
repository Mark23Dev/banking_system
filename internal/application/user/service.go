package user

import (
	"banking_system/internal/domain/user"

	"github.com/google/uuid"
)

type UserService struct {
	repo user.UserRepository
}

func NewUserService(r user.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (u *UserService) AddUser(username, email string, passwordHash, pinHash string) error {
	usr := user.New(username,email, passwordHash, pinHash)
	return u.repo.Save(*usr)
}

// change role to manager (only done by admin user)
func (u *UserService) MakeManager(adminID, userID uuid.UUID) error {
	admin, err := u.repo.FindByID(adminID)
	if err != nil {
		return err
	}
	if !admin.IsAdmin() {
		return user.ErrUnauthorized
	}
	usr, err := u.repo.FindByID(userID)
	if err != nil {
		return err
	}
	usr.ChangeRole(user.Manager)
	return u.repo.Update(usr)
}

func (u *UserService) UpdatePassword(userID uuid.UUID, passwordHash string) error {
	usr, err := u.repo.FindByID(userID)
	if err != nil {
		return err
	}
	usr.ChangePassword(passwordHash)
	return u.repo.Update(usr)
}

func (u *UserService) UpdatePIN(userID uuid.UUID, pinHash string) error {
	usr, err := u.repo.FindByID(userID)
	if err != nil {
		return err
	}
	usr.ChangePIN(pinHash)
	return u.repo.Update(usr)
}

func (u *UserService) CreateCustomerAccount(managerID, userID uuid.UUID) error {
	manager, err := u.repo.FindByID(managerID)
	if err != nil {
		return err
	}
	if !manager.IsManager() {
		return user.ErrUnauthorized
	}

	usr, err := u.repo.FindByID(userID)
	if err != nil {
		return err
	}

	if err := usr.Approve(); err != nil {
		return err
	}

	if err := usr.PromoteToCustomer(); err != nil {
		return err
	}

	return u.repo.Update(usr)
}

