package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}
	req.Header.Set("user-agent", "gator")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}
	defer res.Body.Close()

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(d, &rssFeed); err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}

	return decodeHTML(&rssFeed), nil
}

func decodeHTML(f *RSSFeed) *RSSFeed {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Title)

	for i, item := range f.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		f.Channel.Item[i] = item
	}

	return f
}
