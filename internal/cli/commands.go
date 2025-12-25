package cli

import (
	"fmt"

	"github.com/ElitistNoob/gator/internal/core"
)

type commands struct {
	cmds map[string]func(*core.State, core.Command) error
}

func NewCommand() *commands {
	return &commands{
		make(map[string]func(*core.State, core.Command) error),
	}
}

func (c commands) register(name string, f func(*core.State, core.Command) error) {
	c.cmds[name] = f
}

func (c commands) Run(s *core.State, cmd core.Command) error {
	handler, ok := c.cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}
