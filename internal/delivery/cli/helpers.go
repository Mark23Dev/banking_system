package cli

import (
	"banking_system/internal/domain/account"
	"errors"
	"fmt"
)

func PrintSuccess(msg string) {
	fmt.Println(successStyle.Render("✓ " + msg))
}

func PrintError(msg string) {
	fmt.Println(errorStyle.Render("✗ " + msg))
}

func PrintWarning(msg string) {
	fmt.Println(warningStyle.Render("⚠ " + msg))
}

func PrintInfo(msg string) {
	fmt.Println(infoStyle.Render("ℹ " + msg))
}

func PrintMoney(msg string) {
	fmt.Println(moneyStyle.Render("💰 " + msg))
}

func Divider() {
	fmt.Println("══════════════════════════════════════════════════════════════════════")
}

func PrintBox(title string, body string) {

	fmt.Println(
		boxStyle.Render(
			title + "\n\n" + body,
		),
	)
}

func parseAccountType(input string) (account.AccountType, error) {

	switch input {

	case "1":
		return account.Checking, nil

	case "2":
		return account.Savings, nil

	case "3":
		return account.Business, nil

	case "4":
		return account.Specialty, nil

	default:
		return 0, errors.New("invalid account type")
	}
}