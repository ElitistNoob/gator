package app

import (
	"context"
	"fmt"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *core.State, c core.Command, user db.User) (string, error)) func(s *core.State, c core.Command) (string, error) {
	return func(s *core.State, c core.Command) (string, error) {
		user, err := s.DB.GetUser(context.Background(), s.Cfg.Current_user_name)
		if err != nil {
			return "", fmt.Errorf("couldn't get user: %w", err)
		}

		return handler(s, c, user)
	}
}
