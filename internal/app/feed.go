package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func AddFeed(s *core.State, c core.Command, user db.User) (string, error) {
	if len(c.Args) != 2 {
		return "", fmt.Errorf("expects arguments: <name> <url>\nGot: %v", c.Args)
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
		return "", err
	}

	var str strings.Builder
	fmt.Fprintf(&str, "feed was successfully created:\n")

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
		return "", err
	}

	fmt.Fprintf(&str, "feed followed successfully:\n\n")

	fmt.Fprintf(&str, "> feed_name:   %v\n", feedFollow.FeedName)
	fmt.Fprintf(&str, "> user_name:   %v", feedFollow.UserName)

	return str.String(), nil
}

func GetFeeds(s *core.State, c core.Command) (string, error) {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return "", err
	}

	if len(feeds) < 1 {
		return "", fmt.Errorf("feeds table is empty")
	}

	var str strings.Builder
	fmt.Fprintf(&str, "Found %d feeds:\n\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.DB.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&str, "> ID:           %s\n", feed.ID)
		fmt.Fprintf(&str, "> CreatedAt:    %s\n", feed.CreatedAt)
		fmt.Fprintf(&str, "> UpdateAt:     %s\n", feed.UpdatedAt)
		fmt.Fprintf(&str, "> Name:         %s\n", feed.Name)
		fmt.Fprintf(&str, "> Url:          %v\n", feed.Url)
		fmt.Fprintf(&str, "> Username:     %v", user)
	}

	return str.String(), nil
}
