package root

import (
	"fmt"

	"github.com/ElitistNoob/gator/internal/app"
	"github.com/ElitistNoob/gator/internal/core"
	"github.com/ElitistNoob/gator/internal/tui/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	cmdSelection Mode = iota
	cmdArguments
	cmdOutput
)

type model struct {
	mode            Mode
	state           *core.State
	cursor          int
	commands        []core.Command
	selectedCommand core.Command
	argsInput       []textinput.Model
	argCursor       int
	output          outputMsg
	errMsg          errMsg
}

func ResetModel(m model) model {
	m.mode = cmdSelection
	m.cursor = 0
	m.selectedCommand = core.Command{}
	m.argsInput = nil
	m.argCursor = 0
	m.output = ""
	m.errMsg = errMsg{}

	return m
}

func InitialModel(s *core.State) model {
	return model{
		mode:   cmdSelection,
		state:  s,
		cursor: 0,
		commands: []core.Command{
			{
				Name: "Register User",
				Args: []string{"Name"},
				Run:  wrapCommand(app.RegisterUser),
			},
			{
				Name: "Login",
				Args: []string{"Name"},
				Run:  wrapCommand(app.Login),
			},
			{
				Name: "Display Users",
				Args: []string{},
				Run:  wrapCommand(app.GetUsers),
			},
			{
				Name: "Add New Feed",
				Args: []string{"Name", "Url"},
				Run:  wrapCommand(app.MiddlewareLoggedIn(app.AddFeed)),
			},
			{
				Name: "List Feeds",
				Args: []string{},
				Run:  wrapCommand(app.GetFeeds),
			},
			{
				Name: "Follow Feed",
				Args: []string{"Url"},
				Run:  wrapCommand(app.MiddlewareLoggedIn(app.FollowFeed)),
			},
			{
				Name: "Displays User's Feeds",
				Args: []string{},
				Run:  wrapCommand(app.MiddlewareLoggedIn(app.Following)),
			},
			{
				Name: "Clear DB",
				Args: []string{},
				Run:  wrapCommand(app.ResetDB),
			},
		},
		argCursor: 0,
	}
}

func (m model) selectCommand(idx int) (model, tea.Cmd) {
	m.selectedCommand = m.commands[idx]
	if len(m.selectedCommand.Args) < 1 {
		m.mode = cmdOutput
		return m, runCommandCmd(m)
	}

	m.mode = cmdArguments
	m.argsInput = make([]textinput.Model, len(m.selectedCommand.Args))

	for i, argument := range m.selectedCommand.Args {
		ti := textinput.New()
		ti.Placeholder = fmt.Sprintf("enter %s", argument)
		ti.Width = 22
		ti.Prompt = styles.CursorStyle.Render(fmt.Sprintf("%s: ", argument))
		if i == 0 {
			ti.Focus()
		}
		m.argsInput[i] = ti
	}

	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}
