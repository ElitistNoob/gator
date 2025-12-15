package main

import (
	"context"
	"fmt"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func handler_follow(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("no argument passed\nexpected: <url>\ngot: $v\n", c.args)
	}

	ctx, url := context.Background(), c.args[0]
	feed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't get feed with url: %v\nerr: %w", url, err)
	}

	user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't find user record with current user")
	}

	time := time.Now().UTC()
	args := db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time,
		UpdatedAt: time,
		UserID:    user.ID,
		FeedID:    feed.ID
	}
	d, err := s.db.CreateFeedFollow(ctx, args)
	if err != nil {
		return fmt.Errorf("couldn't create new feed_follow\nerr: %w", err)
	}

	fmt.Printf("feed_follow was successfully created:\n")

	fmt.Printf("> feed_name:           %v\n", feed.name)
	fmt.Printf("> user_name:   %v\n", user.name)

	return nil
}
