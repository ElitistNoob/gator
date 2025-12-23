package main

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/ElitistNoob/gator/internal/database"
)

func handlerBrowse(s *state, c command, user db.User) error {
	limit := 2
	if len(c.args) == 1 {
		arg, err := strconv.Atoi(c.args[0])
		if err != nil {
			return fmt.Errorf("error converting argument to int: %w", err)
		}
		limit = arg
	}

	if limit <= 0 {
		return fmt.Errorf("limit must be greater than 0")
	}

	posts, err := s.db.GetPostsForUser(context.Background(),
		db.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
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
