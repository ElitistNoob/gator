package root

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.errMsg = msg
		return m, nil

	case outputMsg:
		m.output = msg
		m.mode = cmdOutput
		return m, nil

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
			}
		}
	}

	if len(m.argsInput) > 0 {
		ti, cmd := m.argsInput[m.argCursor].Update(msg)
		m.argsInput[m.argCursor] = ti
		return m, cmd
	}
	return m, nil
}
