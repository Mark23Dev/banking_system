package cli

import (
	domainuser "banking_system/internal/domain/user"
	"fmt"
)

func (c *CLI) showDashboard() {

	if !c.session.IsAuthenticated() {
		c.showGuestDashboard()
		return
	}

	switch c.session.CurrentUser().Role {

	case domainuser.Admin:
		c.showAdminDashboard()

	case domainuser.Manager:
		c.showManagerDashboard()

	case domainuser.Customer:
		c.showCustomerDashboard()

	default:
		c.showGuestDashboard()
	}
}



func (c *CLI) showGuestDashboard() {

	Divider()

	fmt.Println(titleStyle.Render("Welcome to DevMak Banking"))

	fmt.Println()

	fmt.Println("Role : Guest")

	fmt.Println()

	fmt.Println("Available Commands")

	fmt.Println("────────────────────────────")

	fmt.Println("signup")
	fmt.Println("login")
	fmt.Println("help")
	fmt.Println("version")
	fmt.Println("exit")

	Divider()
}

func (c *CLI) showCustomerDashboard() {

	user := c.session.CurrentUser()

	Divider()

	fmt.Printf("Welcome back, %s!\n\n", user.Username)

	fmt.Println("Role : Customer")

	fmt.Println()

	fmt.Println("Available Commands")

	fmt.Println("────────────────────────────")

	fmt.Println("profile")
	fmt.Println("accounts")
	fmt.Println("balance")
	fmt.Println("deposit")
	fmt.Println("withdraw")
	fmt.Println("transfer")
	fmt.Println("statement")
	fmt.Println("request-account")
	fmt.Println("logout")

	Divider()
}

func (c *CLI) showManagerDashboard() {

	user := c.session.CurrentUser()

	Divider()

	fmt.Printf("Welcome back, %s!\n\n", user.Username)

	fmt.Println("Role : Manager")

	fmt.Println()

	fmt.Println("Available Commands")

	fmt.Println("────────────────────────────")

	fmt.Println("requests")
	fmt.Println("approve")
	fmt.Println("reject")
	fmt.Println("customers")
	fmt.Println("accounts")
	fmt.Println("logout")

	Divider()
}

func (c *CLI) showAdminDashboard() {

	user := c.session.CurrentUser()

	Divider()

	fmt.Printf("Welcome back, %s!\n\n", user.Username)

	fmt.Println("Role : Administrator")

	fmt.Println()

	fmt.Println("Available Commands")

	fmt.Println("────────────────────────────")

	fmt.Println("users")
	fmt.Println("makemanager")
	fmt.Println("managers")
	fmt.Println("audit")
	fmt.Println("logout")

	Divider()
}