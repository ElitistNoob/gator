package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func getTermSize() (width int, height int) {
	w, h, err := term.GetSize(os.Stdin.Fd())

	maxWidth := 80
	w = w - 2

	if err != nil {
		w = maxWidth
	}

	if w > maxWidth {
		w = maxWidth
	}

	return w, h
}

var termWidth, _ = getTermSize()

var Header = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var Content = lipgloss.NewStyle().
	Width(termWidth).
	Padding(1, 2).
	Margin(0).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#3b3b3b"))

var Error = lipgloss.NewStyle().
	Width(termWidth).
	Padding(1, 2).
	Margin(0).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("9")).
	Foreground(lipgloss.Color("9"))

var Footer = lipgloss.NewStyle().
	Width(termWidth).
	Padding(0, 1).
	Margin(0).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#3b3b3b"))

var Result = lipgloss.NewStyle().
	Width(Content.GetWidth()-6).
	Padding(1, 2).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground((lipgloss.Color("#3b3b3b")))

var Input = lipgloss.NewStyle().
	BorderBottom(true).
	BorderBottomForeground(lipgloss.Color("#3b3b3b")).
	BorderStyle(lipgloss.NormalBorder())

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

var FooterCmdStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("5")).Bold(true)

var Highlight = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

var TitleStyle = lipgloss.NewStyle().
	BorderStyle(func() lipgloss.Border {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return b
	}()).Padding(0, 1)

var InfoStyle = lipgloss.NewStyle().
	BorderStyle(func() lipgloss.Border {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return b
	}())
