package root

import (
	"fmt"
	"strings"

	"github.com/ElitistNoob/gator/internal/tui/styles"
)

func (m model) View() string {
	switch m.mode {
	case cmdArguments:
		return argumentView(m)
	case cmdOutput:
		return outputView(m)
	default:
		return selectionView(m)
	}
}

func selectionView(m model) string {
	var str strings.Builder
	lines := make([]string, 0, len(m.commands))

	lines = append(lines, styles.Header.Render("Select a command to run:\n"))
	for i, command := range m.commands {
		cursor := " "
		if m.cursor == i {
			cursor = "•"
		}

		c := fmt.Sprintf("[%s]", cursor)
		lines = append(lines, fmt.Sprintf("%s %s", styles.CursorStyle.Render(c), command.Name))
	}
	str.WriteString(styles.Content.Render(strings.Join(lines, "\n")))

	footerText := fmt.Sprintf("Press %s to quit.", styles.FooterCmdStyle.Render("q"))
	fmt.Fprintf(&str, "\n%s\n", styles.Footer.Render(footerText))

	return str.String()
}

func argumentView(m model) string {
	var str strings.Builder
	lines := make([]string, 0, len(m.argsInput))

	headerText := styles.Header.Render("Enter arguments:\n")
	lines = append(lines, headerText)

	for _, arg := range m.argsInput {
		lines = append(lines, styles.Input.Render(arg.View()))
	}
	str.WriteString(styles.Content.Render(strings.Join(lines, "\n\n")))

	err := m.errMsg.err
	if err != nil {
		fmt.Fprintf(&str, "\n%s", styles.Error.Render(m.errMsg.Error()))
	}

	var menu strings.Builder
	options := make([]string, 0, len(footerOptions))
	for _, c := range footerOptions {
		options = append(options, fmt.Sprintf("%s %s", styles.FooterCmdStyle.Render(c.Input), c.Action))
	}
	menu.WriteString(strings.Join(options, " "))

	fmt.Fprintf(&str, "\n%s\n", styles.Footer.Render(menu.String()))
	return str.String()
}

func outputView(m model) string {
	var str strings.Builder
	fmt.Fprintf(&str, "%s\n", styles.Content.Render(string(m.output)))

	footerText := fmt.Sprintf("Press %s to return to menu", styles.FooterCmdStyle.Render("esc"))
	fmt.Fprintf(&str, "%s\n", styles.Footer.Render(footerText))

	return str.String()
}

type menuItem struct {
	Input  string
	Action string
}

var footerOptions = []menuItem{
	{
		Input:  "[tab]",
		Action: "next",
	},
	{
		Input:  "[esc]",
		Action: "main menu",
	},
	{
		Input:  "[enter]",
		Action: "run",
	},
	{
		Input:  "[ctrl+c]",
		Action: "quit",
	},
}
