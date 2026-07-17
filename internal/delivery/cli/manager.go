package cli

import (
	"errors"
	"fmt"
	"strconv"
)

func (c *CLI) pendingRequests(args []string) error {

	requests, err := c.requests.PendingRequests()
	if err != nil {
		return err
	}

	if len(requests) == 0 {
		PrintInfo("No pending account requests.")
		return nil
	}

	Divider()

	fmt.Printf("%-4s %-36s %-12s %-20s\n",
		"No", "User ID", "Type", "Created")

	Divider()

	for i, r := range requests {

		fmt.Printf(
			"%-4d %-36s %-12v %-20s\n",
			i+1,
			r.UserID,
			r.AccountType,
			r.CreatedAt.Format("2006-01-02 15:04"),
		)

	}

	Divider()

	return nil
}

func (c *CLI) approveRequest(args []string) error {
	fmt.Println("CLI approveRequest called")
	if len(args) != 1 {
		PrintInfo("Usage: approve <number>")
		return nil
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	requests, err := c.requests.PendingRequests()
	if err != nil {
		return err
	}
	fmt.Println("before")
	fmt.Printf("index=%d, len(requests)=%d\n", index, len(requests))
	if index < 1 || index > len(requests) {
		return errors.New("invalid request number")
	}
	fmt.Println("after")
	
	request := requests[index-1]
	

	manager := c.session.CurrentUser()
	fmt.Println("About to call service")
	if err := c.requests.ApproveRequest(
		manager.ID,
		request.ID,
	); err != nil {
		return err
	}

	PrintSuccess("Request approved.")

	return nil
}

func (c *CLI) rejectRequest(args []string) error {

	if len(args) != 1 {
		PrintInfo("Usage: reject <number>")
		return nil
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	requests, err := c.requests.PendingRequests()
	if err != nil {
		return err
	}

	if index < 1 || index > len(requests) {
		return errors.New("invalid request number")
	}

	selectedRequest := requests[index-1]

	manager := c.session.CurrentUser()

	if err := c.requests.RejectRequest(
		manager.ID,
		selectedRequest.ID,
	); err != nil {
		return err
	}

	PrintSuccess("Request rejected.")

	return nil
}