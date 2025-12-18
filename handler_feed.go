package main

import (
	"context"
	"fmt"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, c command, user db.User) error {
	if len(c.args) != 2 {
		return fmt.Errorf("expects arguments: <name> <url>\nGot: %v", c.args)
	}

	ctx := context.Background()
	name, url := c.args[0], c.args[1]
	args := db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(ctx, args)
	if err != nil {
		return err
	}

	fmt.Printf("feed was successfully created:\n")

	now := time.Now().UTC()
	createFeedFollowArgs := db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(ctx, createFeedFollowArgs)
	if err != nil {
		return err
	}

	fmt.Printf("feed followed successfully:\n")

	fmt.Printf("> feed_name:   %v\n", feedFollow.FeedName)
	fmt.Printf("> user_name:   %v\n", feedFollow.UserName)

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

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("> ID:           %s\n", feed.ID)
		fmt.Printf("> CreatedAt:    %s\n", feed.CreatedAt)
		fmt.Printf("> UpdateAt:     %s\n", feed.UpdatedAt)
		fmt.Printf("> Name:         %s\n", feed.Name)
		fmt.Printf("> Url:          %v\n", feed.Url)
		fmt.Printf("> Username:     %v\n", user)
		fmt.Println(" ")
		fmt.Println("====================================")
	}

	return nil
}
