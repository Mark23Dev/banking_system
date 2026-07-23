package cli

import (
	"banking_system/internal/domain/account"
	"errors"
	"fmt"
	"strconv"
)

func (c *CLI) accounts(args []string) error {

	user := c.session.CurrentUser()

	if user == nil {
		return errors.New("please login first")
	}

	accts, err := c.accts.AccountsByCustomer(user.ID)
	if err != nil {
		return err
	}

	if len(accts) == 0 {
		PrintWarning("You have no bank accounts.")
		return nil
	}

	Divider()

	fmt.Printf(
		"%-15s %-12s %-12s %-10s\n",
		"Account No",
		"Type",
		"Status",
		"Balance",
	)

	Divider()

	for _, a := range accts {

		fmt.Printf(
			"%-15s %-12v %-12v KES %d\n",
			a.AccountNumber,
			a.AccountType,
			a.Status,
			a.Balance,
		)
	}

	Divider()

	return nil
}

func (c *CLI) balance(args []string) error {

	currentUser := c.session.CurrentUser()

	accts, err := c.accts.AccountsByCustomer(currentUser.ID)
	if err != nil {
		return err
	}

	if len(accts) == 0 {
		PrintWarning("You do not have any bank accounts.")
		return nil
	}

	var selected account.Account

	if len(accts) == 1 {
		selected = accts[0]
	} else {

		fmt.Println()
		PrintInfo("Select an account")
		Divider()

		for i, acc := range accts {
			fmt.Printf(
				"%d. %s (%v)\n",
				i+1,
				acc.AccountNumber,
				acc.AccountType,
			)
		}

		choice, err := c.Prompt("Account: ")
		if err != nil {
			return err
		}

		index, err := strconv.Atoi(choice)
		if err != nil {
			return err
		}

		if index < 1 || index > len(accts) {
			return errors.New("invalid account selection")
		}

		selected = accts[index-1]
	}

	Divider()

	fmt.Printf("Account Number : %s\n", selected.AccountNumber)
	fmt.Printf("Account Type   : %v\n", selected.AccountType)
	fmt.Printf("Balance        : KES %d\n", selected.Balance)

	Divider()

	return nil
}