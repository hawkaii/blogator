package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"
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
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %v <url>", cmd.Name)
		os.Exit(1)
	}

	time_between_reqs_string := cmd.Args[0]
	time_between_reqs, err := time.ParseDuration(time_between_reqs_string)
	if err != nil {
		fmt.Println("Invalid time duration")
		os.Exit(1)
	}

	fmt.Println("Collecting feeds every", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func printRSSItem(item RSSItem) {
	fmt.Println("Title:", item.Title)
	fmt.Println("Link:", item.Link)
	fmt.Println("Description:", item.Description)
	fmt.Println("PubDate:", item.PubDate)
	fmt.Println()
}
