package cli


import "github.com/charmbracelet/lipgloss"

var (

	// Titles
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("39"))

	// Success
	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("42")).
		Bold(true)

	// Errors
	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true)

	// Warnings
	warningStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("220")).
		Bold(true)

	// Information
	infoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("45"))

	// Money
	moneyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("46")).
		Bold(true)

	// Prompt
	promptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("69")).
		Bold(true)

	// Boxes
	boxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("69")).
		Padding(1, 2)
)