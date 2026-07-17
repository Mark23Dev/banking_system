package main

import (
	"banking_system/internal/application/account"
	"banking_system/internal/application/accountrequest"
	"banking_system/internal/application/auth"
	"banking_system/internal/application/session"
	"banking_system/internal/application/transaction"
	"banking_system/internal/application/user"
	"banking_system/internal/bootstrap"
	"banking_system/internal/config"
	"banking_system/internal/delivery/cli"
	accountrequeststorage "banking_system/internal/storage/filestorage/accountrequest"
	accountstorage "banking_system/internal/storage/filestorage/accounts"
	transactionstorage "banking_system/internal/storage/filestorage/transactions"
	userstorage "banking_system/internal/storage/filestorage/users"
	"log"
)

func main() {
	cfg := config.Load()

	// storage
	userRepo := userstorage.NewFileUserStore(cfg.UsersFile)
		if err := bootstrap.SeedAdmin(userRepo); err != nil {
		log.Fatal(err)
	}
	accountRepo := accountstorage.NewFileAccountsStore(cfg.AccountsFile)
	accountRequestRepo := accountrequeststorage.NewFileAccountRequestsStore(cfg.AccountRequestsFile)
	transactionRepo := transactionstorage.NewFileTransactionsStore(cfg.TransactionsFile)

	// services
	authService := auth.NewAuthService(userRepo)
	userService := user.NewUserService(userRepo)
	accountService := account.NewAccountService(accountRepo)
	transactionService := transaction.NewTransactionService(transactionRepo, accountRepo)
	accountRequestService := accountrequest.NewAccountRequestService(accountRequestRepo, userRepo, accountRepo)

	sessionManager := session.NewManager()

	app, err := cli.New(
		authService,
		userService,
		accountService,
		accountRequestService,
		transactionService,
		sessionManager,
	)
	if err != nil {
		panic(err)
	}

	app.Run()
}