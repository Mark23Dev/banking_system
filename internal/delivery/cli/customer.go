package cli

import (
	"fmt"
	"strconv"
)

func (c *CLI) deposit(args []string) error {
	amountStr, err := c.Prompt("Enter amount: ")
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	accountNumber, err := c.Prompt("Enter account number: ")
	if err != nil {
		return err
	}

	return c.accts.DepositToAccount(accountNumber, amount)
}

func (c *CLI) withdraw(args []string) error {
	amountStr, err := c.Prompt("Enter amount: ")
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	accountNumber, err := c.Prompt("Enter account number: ")
	if err != nil {
		return err
	}

	return c.accts.WithdrawFromAccount(accountNumber, amount)
}

func (c *CLI) transfer(args []string) error {
	fromAccount, err := c.Prompt("From account: ")
	if err != nil {
		return err
	}

	toAccount, err := c.Prompt("To account: ")
	if err != nil {
		return err
	}

	amountStr, err := c.Prompt("Amount: ")
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	description, err := c.Prompt("Description (optional): ")
if err != nil {
	return err
}

if description == "" {
	description = "Transfer"
}

return c.transactions.Transfer(
	fromAccount,
	toAccount,
	amount,
	description,
)
}

func (c *CLI) statement(args []string) error {
	accountNumber, err := c.Prompt("Enter account number: ")
	if err != nil {
		return err
	}

	acc, err := c.accts.Statement(accountNumber)
	if err != nil {
		return err
	}

	Divider()
	fmt.Printf("Account Number : %s\n", acc.AccountNumber)
	fmt.Printf("Owner          : %s\n", acc.CustomerID)
	fmt.Printf("Type           : %s\n", acc.AccountType)
	fmt.Printf("Balance        : %d\n", acc.Balance)
	fmt.Printf("Status         : %s\n", acc.Status)
	fmt.Printf("Created        : %s\n", acc.CreatedAt.Format("2006-01-02 15:04"))
	Divider()

	return nil
}