package cli

import (
	"fmt"
	"os"
)

func (c *CLI) help(args []string) error {

	Divider()

	fmt.Println(titleStyle.Render("COMMANDS"))

	fmt.Println()

	fmt.Println("Authentication")
	fmt.Println("────────────────────────────────────")
	fmt.Println("signup              Create a user account")
	fmt.Println("login               Login")
	fmt.Println("logout              Logout")

	fmt.Println()

	fmt.Println("Accounts")
	fmt.Println("────────────────────────────────────")
	fmt.Println("accounts            View accounts")
	fmt.Println("balance             View account balance")
	fmt.Println("deposit             Deposit money")
	fmt.Println("withdraw            Withdraw money")
	fmt.Println("transfer            Transfer money")
	fmt.Println("statement           View statement")

	fmt.Println()

	fmt.Println("Requests")
	fmt.Println("────────────────────────────────────")
	fmt.Println("request-account     Request a new bank account")
	fmt.Println("notifications       View notifications")

	fmt.Println()

	fmt.Println("Administration")
	fmt.Println("────────────────────────────────────")
	fmt.Println("approve             Approve account request")
	fmt.Println("reject              Reject account request")
	// fmt.Println("users               View users")
	fmt.Println("makemanager         Promote user to manager")

	fmt.Println()

	fmt.Println("General")
	fmt.Println("────────────────────────────────────")
	fmt.Println("help                Show this help page")
	fmt.Println("clear               Clear the screen")
	fmt.Println("version             Show version")
	fmt.Println("exit                Exit the application")

	Divider()

	return nil
}

func (c *CLI) version(args []string) error {

	Divider()

	PrintInfo("DevMak Banking System")
	PrintInfo("Version : v1.0.0")
	PrintInfo("Build   : 2026.07.15")
	PrintInfo("Go      : 1.25")

	Divider()

	return nil
}

func (c *CLI) clear(args []string) error {

	fmt.Print("\033[H\033[2J")

	return nil
}

func (c *CLI) exit(args []string) error {

	ExitBanner()

	c.rl.Close()

	os.Exit(0)

	return nil
}
