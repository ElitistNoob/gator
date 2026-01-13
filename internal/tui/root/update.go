package root

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	headerHeight := lipgloss.Height(m.outputHeader())
	footerHeight := lipgloss.Height(m.outputFooter())
	menuHeight := lipgloss.Height(m.footer())
	verticalMarginHeight := headerHeight + footerHeight + menuHeight + 4

	switch msg := msg.(type) {

	case errMsg:
		m.errMsg = msg
		return m, nil

	case outputMsg:
		m.output = msg
		m.mode = cmdOutput

		if m.ready {
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(string(m.output))
		}

		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		key := msg.String()
		switch m.mode {
		case cmdSelection:
			switch key {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.commands)-1 {
					m.cursor++
				}
			case "enter", " ":
				m, c := m.selectCommand(m.cursor)
				return m, c
			}

		case cmdArguments:
			switch key {
			case "tab":
				if len(m.argsInput) > 0 {
					m.argsInput[m.argCursor].Blur()
					m.argCursor = (m.argCursor + 1) % len(m.argsInput)
					m.argsInput[m.argCursor].Focus()
				}
				return m, nil
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				return m, runCommandCmd(m)
			case "esc":
				m = ResetModel(m)
				return m, nil
			}

		case cmdOutput:
			switch key {
			case "esc":
				m = ResetModel(m)
				return m, nil
			case "up", "k":
				m.viewport.ScrollUp(1)
				return m, nil
			case "down", "j":
				m.viewport.ScrollDown(1)
				return m, nil
			case "ctrl+k":
				m.viewport.HalfPageUp()
				return m, nil
			case "ctrl+j":
				m.viewport.HalfPageDown()
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	if len(m.argsInput) > 0 {
		ti, cmd := m.argsInput[m.argCursor].Update(msg)
		m.argsInput[m.argCursor] = ti
		return m, cmd
	}
	return m, nil
}
