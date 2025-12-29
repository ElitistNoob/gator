package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/google/uuid"
)

func FollowFeed(s *core.State, c core.Command, user db.User) (string, error) {
	if len(c.Args) < 1 {
		return "", errors.New("no argument passed\nexpected: <url>")
	}

	ctx, url := context.Background(), c.Args[0]
	feed, err := s.DB.GetFeedByUrl(ctx, url)
	if err != nil {
		return "", fmt.Errorf("couldn't get feed with url: %v\nerr: %w", url, err)
	}

	now := time.Now().UTC()
	args := db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	d, err := s.DB.CreateFeedFollow(ctx, args)
	if err != nil {
		return "", err
	}

	var str strings.Builder
	fmt.Fprintf(&str, "feed followed successfully:\n")

	fmt.Fprintf(&str, "> feed_name:   %v\n", d.FeedName)
	fmt.Fprintf(&str, "> user_name:   %v\n", d.UserName)

	return str.String(), nil
}

func Following(s *core.State, c core.Command, user db.User) (string, error) {
	ctx := context.Background()

	feeds, err := s.DB.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return "", err
	}

	var str strings.Builder
	fmt.Fprintf(&str, "following feed:\n")
	for _, feed := range feeds {
		fmt.Fprintf(&str, "> ID:   %v\n", feed.ID)
		fmt.Fprintf(&str, "> CreateAt:   %v\n", feed.CreatedAt)
		fmt.Fprintf(&str, "> UpdatedAt:   %v\n", feed.UpdatedAt)
		fmt.Fprintf(&str, "> feed_name:   %v\n", feed.FeedName)
		fmt.Fprintf(&str, "> user_name:   %v\n", feed.UserName)
	}

	return str.String(), nil
}

func Unfollow(s *core.State, c core.Command, user db.User) (string, error) {
	if len(c.Args) != 1 {
		return "", fmt.Errorf("expected: %s <feed_url>\ngot: %s <null>", c.Name, c.Name)
	}

	ctx := context.Background()
	feed, err := s.DB.GetFeedByUrl(ctx, c.Args[0])
	if err != nil {
		return "", fmt.Errorf("couldn't get feed:\n err: %w", err)
	}

	args := db.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.DB.DeleteFeedFollow(context.Background(), args); err != nil {
		return "", fmt.Errorf("couldn't delete follow record: %w", err)
	}

	str := fmt.Sprintf("%s unfollowed successfully", feed.Name)
	return str, nil
}
