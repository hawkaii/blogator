package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, url string) (*RSSFeed, error) {
	//https newReqeustwithContext
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't create request: %w", err)
	}

	//https defaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	//set the user-agent header to gator with req.Header.Add
	req.Header.Add("User-Agent", "gator")

	//io.ReadAll

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed) //https xml.Unmarshal
	if err != nil {
		return nil, fmt.Errorf("couldn't decode response: %w", err)
	}

	// use html.UnescapeString to decode the feed title
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	return &feed, nil
}

func handler_agg(s *state, cmd command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %v <url>", cmd.Name)
	// }
	//
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n", item.Title)
		fmt.Printf("  %s\n", item.Link)
		fmt.Printf("  %s\n", item.PubDate)
		fmt.Printf("  %s\n", item.Description)
		fmt.Println()
	}

	return nil
}
