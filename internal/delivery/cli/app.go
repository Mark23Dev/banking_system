package cli

import (
	appaccount "banking_system/internal/application/account"
	appaccountrequest "banking_system/internal/application/accountrequest"
	appauth "banking_system/internal/application/auth"
	appsession "banking_system/internal/application/session"
	apptransaction "banking_system/internal/application/transaction"
	appuser "banking_system/internal/application/user"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

type CLI struct {
	rl *readline.Instance

	auth         *appauth.AuthService
	users        *appuser.UserService
	accts      *appaccount.AccountService
	requests      *appaccountrequest.AccountRequestService
	transactions  *apptransaction.TransactionService

	session *appsession.Manager

	commands map[string]CommandFunc
}

func New(
	auth *appauth.AuthService,
	users *appuser.UserService,
	accounts *appaccount.AccountService,
	requests *appaccountrequest.AccountRequestService,
	transactions *apptransaction.TransactionService,
	session *appsession.Manager,
) (*CLI, error) {

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "guest@devmak-bank > ",
		HistoryFile:     "/tmp/devmak-bank.history",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		return nil, err
	}

	cli := &CLI{
		rl: rl,

		auth: auth,
		users: users,
		accts: accounts,
		requests: requests,
		transactions: transactions,

		session: session,
	}

	cli.registerCommands()

	return cli, nil
}


func (c *CLI) Run() {

	defer c.rl.Close()

	c.startup()

	for {

		c.updatePrompt()

		command, err := c.rl.Readline()

		if err != nil {
			break
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		if err := c.route(command); err != nil {
			c.printError(err)
		}
	}
}

func (c *CLI) startup() {

	ShowBanner()

	Loading("Loading configuration...")
	PrintSuccess("Configuration loaded")

	Loading("Loading repositories...")
	PrintSuccess("Repositories loaded")

	Loading("Loading authentication...")
	PrintSuccess("Authentication loaded")

	Loading("Loading notifications...")
	PrintSuccess("Notification service loaded")

	Divider()

	PrintInfo(`Type "help" to view available commands.`)

	fmt.Println()

	c.showGuestDashboard()
}

func (c *CLI) printError(err error) {
	PrintError(err.Error())
}

