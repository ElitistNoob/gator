package root

import (
	"context"

	"github.com/ElitistNoob/gator/internal/core"
	tea "github.com/charmbracelet/bubbletea"
)

func wrapCommand(handler func(*core.State, core.Command) (string, error)) func(ctx context.Context, state *core.State, args []string) (string, error) {
	return func(ctx context.Context, state *core.State, args []string) (string, error) {
		cmd := core.Command{Args: args}
		return handler(state, cmd)
	}
}

func collectArgs(m model) []string {
	args := make([]string, len(m.argsInput))
	for i, ti := range m.argsInput {
		args[i] = ti.Value()
	}
	return args
}

func runCommandCmd(m model) tea.Cmd {
	return func() tea.Msg {
		out, err := m.selectedCommand.Run(context.Background(), m.state, collectArgs(m))
		if err != nil {
			return errMsg{err}
		}

		return outputMsg(out)
	}
}

type outputMsg string
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }
