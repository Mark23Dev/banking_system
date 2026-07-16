package cli

import (
	domainuser "banking_system/internal/domain/user"
	"strings"
)

func (c *CLI) Prompt(label string) (string, error) {

	old := c.rl.Config.Prompt

	c.rl.SetPrompt(label)

	input, err := c.rl.Readline()

	c.rl.SetPrompt(old)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func (c *CLI) PromptPassword(label string) (string, error) {

	pass, err := c.rl.ReadPassword(label)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(pass)), nil
}

func (c *CLI) updatePrompt() {

	if !c.session.IsAuthenticated() {

		c.rl.SetPrompt(promptStyle.Render("guest@devmak-bank > "))

		return
	}

	u := c.session.CurrentUser()

	switch u.Role {

	case domainuser.Admin:
		c.rl.SetPrompt(promptStyle.Render("admin@devmak-bank > "))

	case domainuser.Manager:
		c.rl.SetPrompt(promptStyle.Render("manager@devmak-bank > "))

	case domainuser.Customer:
		c.rl.SetPrompt(
			promptStyle.Render(u.Username + "@devmak-bank > "),
		)

	default:
		c.rl.SetPrompt(promptStyle.Render("guest@devmak-bank > "))
	}
}