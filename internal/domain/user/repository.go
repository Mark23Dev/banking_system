package user

import "github.com/google/uuid"

type UserRepository interface {
	Save(user User) error
	FindAll() ([]User, error)
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	Update(updated User) error
}