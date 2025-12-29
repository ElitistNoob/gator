package styles

import (
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func getTermSize() (width int, height int) {
	w, h, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		log.Fatal(err)
	}

	return w, h
}

var termWidth, _ = getTermSize()

var Header = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12")).MarginTop(1)

var Content = lipgloss.NewStyle().
	Width(termWidth/2).
	Padding(1, 2).
	Margin(0).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#3b3b3b"))

var Footer = lipgloss.NewStyle().
	Width(termWidth/2).
	Padding(0, 1).
	Margin(0).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#3b3b3b"))

var Input = lipgloss.NewStyle().
	BorderBottom(true).
	BorderBottomForeground(lipgloss.Color("#3b3b3b")).
	BorderStyle(lipgloss.NormalBorder())

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).Bold(true)

var FooterCmdStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("5")).Bold(true)

var Highlight = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).Bold(true)
