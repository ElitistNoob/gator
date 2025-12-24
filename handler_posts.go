package main

import (
	"context"
	"flag"
	"fmt"

	db "github.com/ElitistNoob/gator/internal/database"
)

func handlerBrowse(s *state, c command, user db.User) error {
	limit, order := 2, "desc"

	f := flag.NewFlagSet("browse", flag.ExitOnError)
	f.IntVar(&limit, "limit", 2, "number of post to show")
	f.StringVar(&order, "order", "desc", "order of post to show by created at")
	if err := f.Parse(c.args); err != nil {
		return err
	}

	if limit <= 0 {
		return fmt.Errorf("limit must be greater than 0")
	}

	posts, err := s.db.GetPostsForUser(context.Background(),
		db.GetPostsForUserParams{
			UserID:  user.ID,
			Column2: order,
			Limit:   int32(limit),
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
