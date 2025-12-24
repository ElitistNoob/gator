package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/ElitistNoob/gator/internal/timeutils"
)

func handlerBrowse(s *state, c command, user db.User) error {
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
	if err := f.Parse(c.args); err != nil {
		return err
	}

	if limit <= 0 {
		return fmt.Errorf("limit must be greater than 0")
	}

	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		return fmt.Errorf("order must be 'asc' or 'desc'")
	}

	var fromDate time.Time
	if fromStr != "" {
		t, err := timeutils.ParseTime(fromStr)
		if err != nil {
			return fmt.Errorf("wrong date format: %w", err)
		}

		fromDate = t
	}

	toDate, err := timeutils.ParseTime(toStr)
	if err != nil {
		return fmt.Errorf("wrong date format: %w", err)
	}

	ctx := context.Background()
	posts, err := s.db.GetPostsForUser(ctx,
		db.GetPostsForUserParams{
			UserID:        user.ID,
			PublishedAt:   fromDate,
			PublishedAt_2: toDate,
			Column4:       order,
			Limit:         int32(limit),
		})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s\n", len(posts), user.Name)
	for i, post := range posts {
		fmt.Println()
		fmt.Printf("> Blog %d:\n", i+1)
		fmt.Printf(" - From: %s\n", post.FeedName)
		fmt.Printf(" - Published on: %s\n", post.PublishedAt.Format("Mon Jan 2"))
		fmt.Printf(" - Title: %s\n", post.Title)
		fmt.Printf(" - Description: %s\n", post.Description.String)
		fmt.Printf(" - Url: %s\n", post.Url)
		fmt.Println()
		fmt.Println("=============================================")
		fmt.Println()
	}

	return nil
}
