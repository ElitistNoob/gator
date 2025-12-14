package main

import (
	"context"
	"fmt"
)

func agg(s *state, c command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()
	f, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't fetch from %v: %w", url, err)
	}
	fmt.Printf("feed: %v\n", f)
	return nil

}
