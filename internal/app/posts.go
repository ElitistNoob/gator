package app

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/ElitistNoob/gator/internal/core"
	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/ElitistNoob/gator/internal/timeutils"
)

func BrowsePosts(s *core.State, c core.Command, user db.User) (string, error) {
	now := time.Now()
	today := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	).Format("2006-01-02")

	var limit int
	var order, fromStr, toStr string

	f := flag.NewFlagSet("browse", flag.ExitOnError)
	f.IntVar(&limit, "limit", 5, "number of post to show")
	f.StringVar(&fromStr, "from", "", "start date (inclusive), format: YYYY-MM-DD")
	f.StringVar(&toStr, "to", today, "end date (exclusive), format: YYYY-MM-DD")
	f.StringVar(&order, "order", "desc", "order of post to show by created at")
	if err := f.Parse(c.Args); err != nil {
		return "", err
	}

	if limit <= 0 {
		return "", fmt.Errorf("limit must be greater than 0")
	}

	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		return "", fmt.Errorf("order must be 'asc' or 'desc'")
	}

	var fromDate time.Time
	if fromStr != "" {
		t, err := timeutils.ParseTime(fromStr)
		if err != nil {
			return "", fmt.Errorf("wrong date format: %w", err)
		}

		fromDate = t
	}

	toDate, err := timeutils.ParseTime(toStr)
	if err != nil {
		return "", fmt.Errorf("wrong date format: %w", err)
	}

	ctx := context.Background()
	posts, err := s.DB.GetPostsForUser(ctx,
		db.GetPostsForUserParams{
			UserID:        user.ID,
			PublishedAt:   fromDate,
			PublishedAt_2: toDate,
			Column4:       order,
			Limit:         int32(limit),
		})
	if err != nil {
		return "", fmt.Errorf("couldn't get posts for user: %w", err)
	}

	var str strings.Builder
	fmt.Fprintf(&str, "Found %d posts for user %s\n", len(posts), user.Name)
	for i, post := range posts {
		fmt.Fprintf(&str, "> Blog %d:\n", i+1)
		fmt.Fprintf(&str, " - From: %s\n", post.FeedName)
		fmt.Fprintf(&str, " - Published on: %s\n", post.PublishedAt.Format("Mon Jan 2"))
		fmt.Fprintf(&str, " - Title: %s\n", post.Title)
		fmt.Fprintf(&str, " - Description: %s\n", post.Description.String)
		fmt.Fprintf(&str, " - Url: %s", post.Url)
	}

	return str.String(), nil
}
