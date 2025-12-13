package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func NewCommand() *commands {
	return &commands{
		make(map[string]func(*state, command) error),
	}
}

func (c commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c commands) Run(s *state, cmd command) error {
	handler, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	return handler(s, cmd)
}
