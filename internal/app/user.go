package app

import (
	"context"
	"errors"
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

	if userName == "" {
		return "", errors.New("please enter a name, name field cannot be empty.")
	}

	args := db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	}

	user, err := s.DB.CreateUser(ctx, args)
	if err != nil {
		return "", fmt.Errorf("%s could not be registered, user already exist", userName)
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return "", err
	}

	var str strings.Builder

	fmt.Fprintf(&str, "%s\n", "user was successfully created:\n")

	fmt.Fprintf(&str, "> ID:           %v\n", user.ID)
	fmt.Fprintf(&str, "> Created_at:   %v\n", user.CreatedAt)
	fmt.Fprintf(&str, "> Updated_at:   %v\n", user.CreatedAt)
	fmt.Fprintf(&str, "> Name:         %s", user.Name)

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
		return "", fmt.Errorf("user does not exist\n\nerr: %w", err)
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return "", err
	}

	return fmt.Sprintf("User %s has been successfully logged in", styles.Highlight.Render(user.Name)), nil
}

func GetUsers(s *core.State, c core.Command) (string, error) {
	currentUser := s.Cfg.Current_user_name

	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting users: %w", err)
	}

	if len(users) == 0 {
		return "", fmt.Errorf("users table is empty")
	}

	var str strings.Builder
	lines := make([]string, 0, len(users))
	usersLen := fmt.Sprintf("%d", len(users))

	headerStr := styles.Header.Render(fmt.Sprintf("%s users found:\n", usersLen))
	lines = append(lines, headerStr)

	for _, user := range users {
		userStr := fmt.Sprintf("%s %s", styles.Highlight.Render("*"), user.Name)
		if user.Name == currentUser {
			userStr += styles.Highlight.Render(" (current)")
		}
		lines = append(lines, userStr)
	}
	str.WriteString(strings.Join(lines, "\n"))

	return str.String(), nil
}
