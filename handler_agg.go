package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func agg(s *state, c command) error {
	if len(c.args) != 1 {
		return fmt.Errorf("usage: agg <timeBetweenReqs>")
	}

	timeBetweenReqs := c.args[0]
	timeInterval, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return fmt.Errorf("invalid duration %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeInterval)

	ctx := context.Background()
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s, ctx); err != nil {
			log.Printf("scrape error: %v", err)
		}
	}
}

func scrapeFeeds(ctx context.Context, s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch next feed: %w", err)
	}

	rssFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed %s: %w", nextFeed.Name, err)
	}

	if err := s.db.MarkFeedFetched(ctx, nextFeed.ID); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}

	var feedTitle string
	for i, item := range rssFeed.Channel.Item {
		if item.Title == "" {
			continue
		}

		if rssFeed.Channel.Title != feedTitle {
			feedTitle = rssFeed.Channel.Title
			fmt.Println()
			fmt.Printf("Feed: %s\n", feedTitle)
		}

		fmt.Printf("> Blog %d: %s\n", i+1, item.Title)
	}

	return nil
}
