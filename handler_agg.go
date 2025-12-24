package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/ElitistNoob/gator/internal/database"
	"github.com/ElitistNoob/gator/internal/dbutils"
	"github.com/ElitistNoob/gator/internal/timeutils"
	"github.com/google/uuid"
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
		if err := scrapeFeeds(ctx, s); err != nil {
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

	fmt.Printf("length: %d\n", len(rssFeed.Channel.Item))
	for _, item := range rssFeed.Channel.Item {
		pubDate, err := timeutils.ParseTime(item.PubDate)
		if err != nil {
			return err
		}

		params := db.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: dbutils.ToNullString(item.Description),
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}

		if _, err := s.db.CreatePost(ctx, params); err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			} else {
				return fmt.Errorf("couldn't create post: %w", err)
			}
		}
	}

	return nil
}
