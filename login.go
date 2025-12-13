package main

import (
	"fmt"
)

func handlerLogin(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("user name is required")
	}

	if err := s.cfg.SetUser(c.args[0]); err != nil {
		return err
	}

	fmt.Printf("User %s has been successfully logged in", c.args[0])
	return nil
}
