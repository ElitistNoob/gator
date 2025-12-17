package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, c command) error {
	if len(c.args) < 1 {
		return errors.New("no argument passed\nexpected: <url>")
	}

	ctx, url := context.Background(), c.args[0]
	feed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't get feed with url: %v\nerr: %w", url, err)
	}

	user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	args := db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	d, err := s.db.CreateFeedFollow(ctx, args)
	if err != nil {
		return err
	}

	fmt.Printf("feed_follow was successfully created:\n")

	fmt.Printf("> feed_name:   %v\n", d.FeedName)
	fmt.Printf("> user_name:   %v\n", d.UserName)

	return nil
}

func handlerFollowing(s *state, c command) error {
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.Current_user_name)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("following feed:\n")
	for _, feed := range feeds {

		fmt.Printf("> ID:   %v\n", feed.ID)
		fmt.Printf("> CreateAt:   %v\n", feed.CreatedAt)
		fmt.Printf("> UpdatedAt:   %v\n", feed.UpdatedAt)
		fmt.Printf("> feed_name:   %v\n", feed.FeedName)
		fmt.Printf("> user_name:   %v\n", feed.UserName)
	}

	return nil
}
