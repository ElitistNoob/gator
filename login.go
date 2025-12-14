package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("user name is required")
	}

	ctx := context.Background()
	userName := c.args[0]

	user, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("user does not exist: %w", err)
		os.Exit(1)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s has been successfully logged in", user.Name)
	return nil
}
