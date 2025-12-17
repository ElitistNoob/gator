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
		return err
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
		return err
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

func handlerGetFeeds(s *state, c command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		return fmt.Errorf("feeds table is empty")
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Println(" ")
		fmt.Printf("> Name:         %s\n", feed.Name)
		fmt.Printf("> Url:          %v\n", feed.Url)
		fmt.Printf("> Username:     %v\n", user)
		fmt.Println(" ")
		fmt.Println("====================================")
	}

	return nil
}
