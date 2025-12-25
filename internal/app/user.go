package app

import (
	"context"
	"fmt"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func RegisterUser(s *core.State, c core.Command) error {
	if len(c.Args) < 1 {
		return fmt.Errorf("a name was not provided")
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
		return err
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("user was successfully created:\n")

	fmt.Printf("> ID:           %v\n", user.ID)
	fmt.Printf("> Created_at:   %v\n", user.CreatedAt)
	fmt.Printf("> Updated_at:   %v\n", user.CreatedAt)
	fmt.Printf("> Name:         %s\n", user.Name)

	return nil
}

func Login(s *core.State, c core.Command) error {
	if len(c.Args) < 1 {
		return fmt.Errorf("user name is required")
	}

	ctx := context.Background()
	userName := c.Args[0]

	user, err := s.DB.GetUser(ctx, userName)
	if err != nil {
		return err
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s has been successfully logged in", user.Name)
	return nil
}

func GetUsers(s *core.State, c core.Command) error {
	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	if len(users) == 0 {
		return fmt.Errorf("users table is empty")
	}

	currentUser := s.Cfg.Current_user_name
	for _, user := range users {
		string := fmt.Sprintf("* %s", user.Name)
		if user.Name == currentUser {
			string += " (current)"
		}

		fmt.Println(string)
	}

	return nil
}
