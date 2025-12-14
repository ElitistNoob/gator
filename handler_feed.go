package main

import (
	"context"
	"fmt"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, c command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("failed getting current user: %v\nerr: %w", s.cfg.Current_user_name, err)
	}

	if len(c.args) < 2 {
		return fmt.Errorf("not enough arguments passed.\nExpects <name> <url>\nGot: %v", c.args)
	}

	name, url := c.args[0], c.args[1]
	args := db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("failed creating feed: %w", err)
	}

	fmt.Printf("feed was successfully created:\n")

	fmt.Printf("> ID:           %v\n", feed.ID)
	fmt.Printf("> Created_at:   %v\n", feed.CreatedAt)
	fmt.Printf("> Updated_at:   %v\n", feed.CreatedAt)
	fmt.Printf("> Name:         %s\n", feed.Name)
	fmt.Printf("> Url:          %v\n", feed.Url)
	fmt.Printf("> UserID:       %v\n", feed.UserID)

	return nil
}
