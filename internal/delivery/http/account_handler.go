package http

import (
	appaccount "banking_system/internal/application/account"
	domainaccount "banking_system/internal/domain/account"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type AccountHandler struct {
	service *appaccount.AccountService
}

type CreateAccountRequest struct {
	CustomerID  uuid.UUID                 `json:"customer_id"`
	AccountType domainaccount.AccountType `json:"account_type"`
}

type DepositRequest struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
}

type WithdrawRequest struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
}

func NewAccountHandler(
	service *appaccount.AccountService,
) *AccountHandler {

	return &AccountHandler{
		service: service,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acct, err := h.service.CreateNewAccount(
		req.CustomerID,
		req.AccountType,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acct)
}

func (h *AccountHandler) GetCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	customerIDStr := r.URL.Query().Get("customer_id")
	if customerIDStr == "" {
		http.Error(w, "customer_id is required", http.StatusBadRequest)
		return
	}

	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		http.Error(w, "invalid customer_id", http.StatusBadRequest)
		return
	}

	accts, err := h.service.AccountsByCustomer(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accts)
}

func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var req DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.DepositToAccount(req.AccountNumber, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deposit successful",
	})
}

func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var req WithdrawRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.WithdrawFromAccount(req.AccountNumber, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Withdawal successful",
	})

}
