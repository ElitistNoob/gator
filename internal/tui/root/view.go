package root

import (
	"fmt"
	"strings"

	"github.com/ElitistNoob/gator/internal/tui/styles"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if !m.ready {
		return ""
	}

	var body string

	switch m.mode {
	case cmdArguments:
		body = argumentView(m)
	case cmdOutput:
		body = fmt.Sprintf("%s\n\n%s\n%s", m.outputHeader(), outputView(m), m.outputFooter())
	default:
		body = selectionView(m)
	}

	return body + "\n" + m.footer()
}

func (m model) outputHeader() string {
	title := styles.TitleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) outputFooter() string {
	info := styles.InfoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m model) footer() string {
	return styles.Footer.Render(
		strings.Join([]string{
			styles.FooterCmdStyle.Render("[↑↓]") + " scroll",
			styles.FooterCmdStyle.Render("[j/k]") + " move",
			styles.FooterCmdStyle.Render("[enter]") + " select/run",
			styles.FooterCmdStyle.Render("[esc]") + " back",
			styles.FooterCmdStyle.Render("[q]") + " quit",
		}, " • "),
	)
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
		lines = append(lines, fmt.Sprintf("%s %s\t", styles.CursorStyle.Render(c), command.Name))
	}
	str.WriteString(styles.Content.Render(strings.Join(lines, "\n")))

	return str.String()
}

func argumentView(m model) string {
	var str strings.Builder
	lines := make([]string, 0, len(m.argsInput))

	headerText := styles.Header.Render(fmt.Sprintf("%s:\n", m.selectedCommand.Name))
	lines = append(lines, headerText)

	for _, arg := range m.argsInput {
		lines = append(lines, styles.Input.Render(arg.View()))
	}
	str.WriteString(styles.Content.Render(strings.Join(lines, "\n\n")))

	err := m.errMsg.err
	if err != nil {
		fmt.Fprintf(&str, "\n%s", styles.Error.Render(m.errMsg.Error()))
	}

	return str.String()
}

func outputView(m model) string {
	var str strings.Builder
	str.WriteString(m.viewport.View())
	str.WriteString("\n")
	return str.String()
}
