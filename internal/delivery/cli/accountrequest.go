package cli

import (
	"banking_system/internal/domain/accountrequest"
	"errors"
	"fmt"
)

func (c *CLI) requestAccount(args []string) error {

	if !c.session.IsAuthenticated() {
		return errors.New("please login first")
	}

	current := c.session.CurrentUser()

	if current.IsCustomer() {
		PrintWarning("You already have a customer account.")
		return nil
	}

	fmt.Println()
	PrintInfo("Create Bank Account")
	fmt.Println()

	fmt.Println("Available account types")
	fmt.Println("-----------------------")
	fmt.Println("1. Checking")
	fmt.Println("2. Savings")
	fmt.Println("3. Business")
	fmt.Println("4. Specialty")

	choice, err := c.Prompt("Choose account type: ")
	if err != nil {
		return err
	}

	accountType, err := parseAccountType(choice)
	if err != nil {
		return err
	}

	request := accountrequest.New(
		current.ID,
		accountType,
	)

	if err := c.requests.Submit(*request.AccountID, request.AccountType); err != nil {
		return err
	}

	PrintSuccess("Account request submitted successfully.")
	PrintInfo("Awaiting manager approval.")

	return nil
}