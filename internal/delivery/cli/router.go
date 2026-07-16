package cli

import (
	"fmt"
	"strings"
)

type CommandFunc func(args []string) error



func (c *CLI) route(input string) error {

	fields := strings.Fields(input)

	if len(fields) == 0 {
		return nil
	}

	command := strings.ToLower(fields[0])

	args := fields[1:]

	handler, ok := c.commands[command]

	if !ok {
		return fmt.Errorf("unknown command: %s", command)
	}

	return handler(args)
}