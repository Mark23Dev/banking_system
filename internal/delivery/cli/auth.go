package cli

func (c *CLI) signup(args []string) error {

	PrintInfo("Create a new user account")

	username, err := c.Prompt("Username: ")
	if err != nil {
		return err
	}

	email, err := c.Prompt("Email: ")
	if err != nil {
		return err
	}

	password, err := c.PromptPassword("Password: ")
	if err != nil {
		return err
	}

	pin, err := c.PromptPassword("PIN: ")
	if err != nil {
		return err
	}

	if err := c.auth.Signup(username, email, password, pin); err != nil {
		return err
	}

	PrintSuccess("User account created successfully.")
	PrintInfo("You can now login.")

	return nil
}

func (c *CLI) login(args []string) error {

	if c.session.IsAuthenticated() {
		PrintWarning("You are already logged in.")
		return nil
	}

	email, err := c.Prompt("Email: ")
	if err != nil {
		return err
	}

	password, err := c.PromptPassword("Password: ")
	if err != nil {
		return err
	}

	u, err := c.auth.Authenticate(email, password)
	if err != nil {
		return err
	}

	c.session.Login(u)

	PrintSuccess("Login successful.")
	PrintInfo("Welcome back, " + u.Username + "!")

	c.showDashboard()

	return nil
}

func (c *CLI) logout(args []string) error {

	if !c.session.IsAuthenticated() {
		PrintWarning("You are not logged in.")
		return nil
	}

	name := c.session.CurrentUser().Username

	c.session.Logout()

	PrintSuccess("Goodbye, " + name + ".")

	return nil
}