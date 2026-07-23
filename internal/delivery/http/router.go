package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(accountHandler *AccountHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.CreateAccount)
		r.Get("/{customer_id}", accountHandler.GetCustomerAccounts)
		r.Put("/deposit", accountHandler.Deposit)
		r.Put("/withdraw", accountHandler.Withdraw)
	})

	return r
}