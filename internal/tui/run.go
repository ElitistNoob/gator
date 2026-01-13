package tui

import (
	"log"

	"github.com/ElitistNoob/gator/internal/app"
	"github.com/ElitistNoob/gator/internal/core"
	"github.com/ElitistNoob/gator/internal/tui/root"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	state, err := app.Initialize(core.ModeTUI)
	if err != nil {
		log.Fatalf("failed to initialize app: %s", err)
	}
	defer state.SQLDB.Close()

	model := root.InitialModel(state)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalf("TUI, failed to run: %v", err)
	}

	return nil
}
