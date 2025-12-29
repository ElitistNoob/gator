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
	fmt.Fprintf(&str, "%s", styles.Header.Render("Select a command to run")+"\n\n")

	var body strings.Builder
	bodyLines := make([]string, 0, len(m.commands))
	for i, command := range m.commands {
		cursor := " "
		if m.cursor == i {
			cursor = "•"
		}

		c := fmt.Sprintf("[%s]", cursor)
		bodyLines = append(bodyLines, fmt.Sprintf("%s %s", styles.CursorStyle.Render(c), command.Name))
	}
	body.WriteString(strings.Join(bodyLines, "\n"))

	fmt.Fprintf(&str, "%s", styles.Content.Render(body.String()))
	footerText := fmt.Sprintf("Press %s to quit.", styles.FooterCmdStyle.Render("q"))
	fmt.Fprintf(&str, "\n%s\n", styles.Footer.Render(footerText))

	return str.String()
}

func argumentView(m model) string {
	var str strings.Builder
	headerText := styles.Header.Render("Enter arguments:")
	fmt.Fprintf(&str, "%s\n\n", headerText)

	var body strings.Builder
	bodyLines := make([]string, 0, len(m.argsInput))
	for _, arg := range m.argsInput {
		bodyLines = append(bodyLines, styles.Input.Render(arg.View()))
	}
	body.WriteString(strings.Join(bodyLines, "\n\n"))
	fmt.Fprintf(&str, "%s\n", styles.Content.Render(body.String()))

	footerOptions := []menuItem{
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

	fmt.Fprintf(&str, "%s\n", styles.Footer.Render(renderHelper(footerOptions)))
	return str.String()
}

func outputView(m model) string {
	var str strings.Builder
	fmt.Fprintf(&str, "\n%s\n", styles.Content.Render(string(m.output)))

	footerText := fmt.Sprintf("Press %s to return to menu", styles.FooterCmdStyle.Render("esc"))
	fmt.Fprintf(&str, "%s\n", styles.Footer.Render(footerText))

	return str.String()
}

type menuItem struct {
	Input  string
	Action string
}

func renderHelper(cmds []menuItem) string {
	var str strings.Builder
	lines := make([]string, 0, len(cmds))
	for _, c := range cmds {
		lines = append(lines, fmt.Sprintf("%s %s", styles.FooterCmdStyle.Render(c.Input), c.Action))
	}
	str.WriteString(strings.Join(lines, " "))

	return str.String()
}
