package main

import (
	"reflect"

	"github.com/ElitistNoob/gator/internal/config"
)

func hasMethod(s *State, cmd string) (reflect.Method, bool) {
	t := reflect.TypeOf(*s)
	return t.MethodByName(cmd)
}

func (c commands) Run(s *State, cmd command) error {
	v := reflect.ValueOf(c)
	m, ok := hasMethod(s, cmd.name)
	if ok {
		m.Func.Call([]reflect.Value{v})
	}

	return nil
}

func (c command) Register(name string, f func(*State, command) error) {
	return nil
}

type State struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*State, command) error
}
