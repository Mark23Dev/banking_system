package cli


func (c *CLI) makeManager(args []string) error {
	username, err := c.Prompt("Username: ")
	if err != nil {
		return err
	}
	target, err := c.users.FindByUsername(username)
	if err != nil {
		return err
	}
	admin := c.session.CurrentUser()

	if err := c.users.MakeManager(admin.ID, target.ID); err != nil {
		return err
	}
	PrintSuccess("user promoted to manager")
	return nil
}