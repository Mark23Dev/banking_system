package bootstrap

import (
	"banking_system/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin(userRepo user.UserRepository) error {
	users, err := userRepo.FindAll()
	if err != nil {
		return err
	}

	// check admin existence
	for _, u := range users {
		if u.IsAdmin() {
			return nil
		}
	}

	// hash default credentials
	passwordHash, err := hashPassword("admin123")
	if err != nil {
		return err
	}

	pinHash, err := hashPassword("1234")
	if err != nil {
		return err
	}

	admin := user.New(
		"admin",
		"admin@devmak.bank",
		passwordHash,
		pinHash,
	)

	// Promote to administrator
	admin.ChangeRole(user.Admin)

	// Activate the account
	if err := admin.Approve(); err != nil {
		return err
	}

	return userRepo.Save(*admin)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}