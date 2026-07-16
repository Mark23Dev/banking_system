package accountrequeststorage

import (
	"banking_system/internal/domain/accountrequest"
	"banking_system/internal/utils"
	"errors"

	"github.com/google/uuid"
)

type FileAccountRequestsStore struct {
	filepath string
}

func NewFileAccountRequestsStore(path string) *FileAccountRequestsStore {
	return &FileAccountRequestsStore{
		filepath: path,
	}
}

func (f *FileAccountRequestsStore) Save(request accountrequest.AccountRequest) error {
	requests, err := f.FindAll()
	if err != nil {
		return err
	}

	requests = append(requests, request)

	return utils.WriteJSON(f.filepath, requests)
}

func (f *FileAccountRequestsStore) FindAll() ([]accountrequest.AccountRequest, error) {
	var requests []accountrequest.AccountRequest
	err := utils.ReadJSON(f.filepath, &requests)
	return requests, err
}

func (f *FileAccountRequestsStore) FindByID(id uuid.UUID) (accountrequest.AccountRequest, error) {
	requests, err := f.FindAll()
	if err != nil {
		return accountrequest.AccountRequest{}, err
	}

	for _, request := range requests {
		if request.ID == id {
			return request, nil
		}
	}

	return accountrequest.AccountRequest{}, errors.New("account request not found")
}

func (f *FileAccountRequestsStore) Update(updated accountrequest.AccountRequest) error {
	requests, err := f.FindAll()
	if err != nil {
		return err
	}

	for i, request := range requests {
		if request.ID == updated.ID {
			requests[i] = updated
			return utils.WriteJSON(f.filepath, requests)
		}
	}

	return errors.New("account request not found")
}