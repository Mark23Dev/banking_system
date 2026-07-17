package cli

func (c *CLI) registerCommands() {
	c.commands = map[string]CommandFunc{
		// System
		"help":    c.help,
		"clear":   c.clear,
		"version": c.version,
		"exit":    c.exit,

		// Authentication
		"signup":  c.signup,
		"login":   c.login,
		"logout":  c.logout,

		// account request
		"request-account": c.requestAccount,

		// manager requests
		"requests": c.pendingRequests,
		"approve":  c.approveRequest,
		"reject":   c.rejectRequest,

		"accounts": c.accounts,
		"balance": c.balance,

		// admin requests
		"makemanager": c.makeManager,

		// customer
		"deposit": c.deposit,
		"withdraw": c.withdraw,
		"transfer": c.transfer,
		"statement": c.statement,
	}
}

