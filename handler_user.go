package main

import (
	"context"
	"fmt"
	"log"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("a name was not provided")
	}

	ctx := context.Background()
	currentTime := time.Now().UTC()
	userName := c.args[0]

	args := db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	}

	user, err := s.db.CreateUser(ctx, args)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user: %s, was successfully created", s.cfg.Current_user_name)
	log.Println(user.Name)
	return nil
}

func handlerLogin(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("user name is required")
	}

	ctx := context.Background()
	userName := c.args[0]

	user, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("user does not exist: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s has been successfully logged in", user.Name)
	return nil
}
