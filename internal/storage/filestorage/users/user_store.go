package userstorage

import (
	"banking_system/internal/domain/user"
	"banking_system/internal/utils"
	"errors"

	"github.com/google/uuid"
)

type FileUserStore struct {
	filepath string
}

func NewFileUserStore(path string) *FileUserStore {
	return &FileUserStore{
		filepath: path,
	}
}

func (f *FileUserStore) Save(u user.User) error {
	users, err := f.FindAll()
	if err != nil {
		return err
	}

	users = append(users, u)

	return utils.WriteJSON(f.filepath, users)
}

func (f *FileUserStore) FindAll() ([]user.User, error) {
	var users []user.User

	err := utils.ReadJSON(f.filepath, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (f *FileUserStore) FindByID(id uuid.UUID) (user.User, error) {
	users, err := f.FindAll()
	if err != nil {
		return user.User{}, err
	}

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}

	return user.User{}, errors.New("user not found")
}

func (f *FileUserStore) FindByEmail(email string) (user.User, error) {
	users, err := f.FindAll()
	if err != nil {
		return user.User{}, err
	}

	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}

	return user.User{}, errors.New("user not found")
}

func (f *FileUserStore) FindByUsername(username string) (user.User, error) {
	users, err := f.FindAll()
	if err != nil {
		return user.User{}, err
	}

	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}

	return user.User{}, errors.New("user not found")
}

func (f *FileUserStore) Update(updated user.User) error {
	users, err := f.FindAll()
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == updated.ID {
			users[i] = updated
			return utils.WriteJSON(f.filepath, users)
		}
	}

	return errors.New("user not found")
}