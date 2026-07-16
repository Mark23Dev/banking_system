package auth

import (
	"banking_system/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}


func (a *AuthService) Authenticate(email string, password string) (*user.User, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	
	if !verifyPassword(password, user.PasswordHash) {
		return nil, ErrInvalidUserCredentials
	} 
	return &user, nil
}

func (a *AuthService) Signup(username, email string, password, pin string) error {
	if _,err := a.repo.FindByEmail(email); err == nil {
		return ErrEmailAlreadyUsed
	}
	if _, err := a.repo.FindByUsername(username); err == nil {
	  return ErrUsernameAlreadyUsed
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}
	pinHash, err := hashPassword(pin)
	if err != nil {
		return err
	}

	newUser := user.New(username, email, passwordHash, pinHash)

	return a.repo.Save(*newUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}