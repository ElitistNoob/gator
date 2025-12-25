package app

import (
	"context"
	"fmt"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func AddFeed(s *core.State, c core.Command, user db.User) error {
	if len(c.Args) != 2 {
		return fmt.Errorf("expects arguments: <name> <url>\nGot: %v", c.Args)
	}

	ctx := context.Background()
	name, url := c.Args[0], c.Args[1]
	args := db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.DB.CreateFeed(ctx, args)
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
	feedFollow, err := s.DB.CreateFeedFollow(ctx, createFeedFollowArgs)
	if err != nil {
		return err
	}

	fmt.Printf("feed followed successfully:\n")

	fmt.Printf("> feed_name:   %v\n", feedFollow.FeedName)
	fmt.Printf("> user_name:   %v\n", feedFollow.UserName)

	return nil
}

func GetFeeds(s *core.State, c core.Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		return fmt.Errorf("feeds table is empty")
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.DB.GetUserById(context.Background(), feed.UserID)
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
