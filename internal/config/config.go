package config


type Config struct {
	UsersFile string
	AccountsFile string
	AccountRequestsFile string
	TransactionsFile string
}

const (
    UsersFile           = "storage/users.json"
    AccountsFile        = "storage/accounts.json"
    AccountRequestsFile = "storage/account_requests.json"
    TransactionsFile    = "storage/transactions.json"
)


func Load() Config {
	return Config{
		UsersFile: UsersFile,
		AccountsFile: AccountsFile,
		AccountRequestsFile: AccountRequestsFile,
		TransactionsFile: TransactionsFile,
	}
}