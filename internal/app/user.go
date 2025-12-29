package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/ElitistNoob/gator/internal/tui/styles"
	"github.com/google/uuid"
)

func RegisterUser(s *core.State, c core.Command) (string, error) {
	if len(c.Args) < 1 {
		return "", fmt.Errorf("a name was not provided")
	}

	ctx := context.Background()
	currentTime := time.Now().UTC()
	userName := c.Args[0]

	args := db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	}

	user, err := s.DB.CreateUser(ctx, args)
	if err != nil {
		return "", err
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return "", err
	}

	var str strings.Builder

	fmt.Fprintf(&str, "%s\n", "user was successfully created:\n")

	fmt.Fprintf(&str, "> ID:           %v\n", user.ID)
	fmt.Fprintf(&str, "> Created_at:   %v\n", user.CreatedAt)
	fmt.Fprintf(&str, "> Updated_at:   %v\n", user.CreatedAt)
	fmt.Fprintf(&str, "> Name:         %s\n", user.Name)

	return str.String(), nil
}

func Login(s *core.State, c core.Command) (string, error) {
	if len(c.Args) < 1 {
		return "", fmt.Errorf("user name is required")
	}

	ctx := context.Background()
	userName := c.Args[0]

	user, err := s.DB.GetUser(ctx, userName)
	if err != nil {
		return "", err
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return "", err
	}

	return fmt.Sprintf("\n\nUser %s has been successfully logged in", styles.Highlight.Render(user.Name)), nil
}

func GetUsers(s *core.State, c core.Command) (string, error) {
	var str strings.Builder

	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting users: %w", err)
	}

	if len(users) == 0 {
		return "no registered users", fmt.Errorf("users table is empty")
	}

	currentUser := s.Cfg.Current_user_name
	fmt.Fprintf(&str, "%d users found\n\n", len(users))
	for i, user := range users {
		fmt.Fprintf(&str, "%s %s", styles.Highlight.Render("*"), user.Name)
		if user.Name == currentUser {
			fmt.Fprintln(&str, " (current)")
		}
		if i != len(users) {
			fmt.Fprintf(&str, "\n")
		}
	}

	return str.String(), nil
}
