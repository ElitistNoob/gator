package main

import (
	"context"
	"fmt"

	db "github.com/ElitistNoob/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, c command, user db.User) error) func(s *state, c command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}

		return handler(s, c, user)
	}
}
