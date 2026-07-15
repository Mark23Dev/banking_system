package accountrequest

import "github.com/google/uuid"

type AccountRequestRepository interface{
	Save(request AccountRequest) error
	FindAll() ([]AccountRequest, error)
	FindByID(id uuid.UUID) (AccountRequest, error)
	Update(updated AccountRequest) error
}
