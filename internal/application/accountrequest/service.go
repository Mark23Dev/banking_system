package accountrequest

import (
	"banking_system/internal/domain/account"
	accountrequest "banking_system/internal/domain/accountrequest"
	"banking_system/internal/domain/user"
	"fmt"

	"github.com/google/uuid"
)

type AccountRequestService struct {
	requests accountrequest.AccountRequestRepository
	users user.UserRepository
	accounts account.AccountRepsitory
}

func NewAccountRequestService(
	requestRepo accountrequest.AccountRequestRepository,
	userRepo user.UserRepository,
	accountRepo account.AccountRepsitory,
) *AccountRequestService {
	return &AccountRequestService{
		requests: requestRepo,
		users: userRepo,
		accounts: accountRepo,
	}
}

func (a *AccountRequestService) ApproveRequest(managerID, requestID uuid.UUID) error {
	fmt.Println("Before approve:")
	manager, err := a.users.FindByID(managerID)
	if err != nil {
		return err
	}
	if !manager.IsManager() {
		return user.ErrUnauthorized
	}
	

	request, err := a.requests.FindByID(requestID)
	if err != nil {
		return err
	}
	
	if err := request.Approve(managerID); err != nil {
		return err
	}
	

	// account creation (manager approval follows creation of an account)
	newAccount, err := account.New(request.UserID, request.AccountType)

	if err != nil {
		return err
	}
	if err := a.accounts.Save(*newAccount); err != nil {
		return err
	}
	request.AccountID = &newAccount.ID
	
	// create customer account for the user
	customer, err := a.users.FindByID(request.UserID)
	if err != nil {
		return err
	}
	customer.Approve()
	if err := customer.PromoteToCustomer(); err != nil {
		return err
	}
	if err := a.users.Update(customer); err != nil {
		return err
	}

	return a.requests.Update(request)
}

func (a *AccountRequestService) RejectRequest(managerID, accountID uuid.UUID) error {
	manager, err := a.users.FindByID(managerID)
	if err != nil {
		return err
	}
	if !manager.IsManager() {
		return user.ErrUnauthorized
	}

	request, err := a.requests.FindByID(accountID)
	if err != nil {
		return err
	}
	if err := request.Reject(managerID); err != nil {
		return err
	}
	return a.requests.Update(request)
}

func (s *AccountRequestService) Submit(
	userID uuid.UUID,
	accountType account.AccountType,
) error {
	usr, err := s.users.FindByID(userID)
	if err != nil {
		return err
	}

	if !usr.IsAccepted() {
		return user.ErrUserNotApproved
	}

	request := accountrequest.New(userID, accountType)

	return s.requests.Save(*request)
}

func (s *AccountRequestService) PendingRequests() ([]accountrequest.AccountRequest, error) {
	var pendingRequests []accountrequest.AccountRequest
	requests, err := s.requests.FindAll()
	if err != nil {
		return nil, err
	}
	for _, req := range requests {
		if req.Status == accountrequest.Pending {
			pendingRequests = append(pendingRequests, req)
		}
	}
	return pendingRequests, nil
}