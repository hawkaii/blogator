package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hawkaii/blogator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 2 {
		fmt.Printf("usage: %v <name> <url>", cmd.Name)
		os.Exit(1)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.dbQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})

	if err != nil {
		fmt.Print("couldn't create feed: %w", err)
		os.Exit(1)
	}

	feed_follow, err := s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		fmt.Print("couldn't create feed follow: %w", err)
		os.Exit(1)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("Followed: details below")
	printFeedFollow(feed_follow.UserName, feed_follow.FeedName)
	fmt.Println()
	fmt.Println("=====================================")

	return nil

}

func handlerGetFeeds(s *state, _ command) error {

	feeds, err := s.dbQueries.GetFeeds(context.Background())

	if err != nil {
		fmt.Print("couldn't get feeds: %w", err)
		os.Exit(1)
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		user, err := s.dbQueries.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Print("couldn't get user: %w", err)
			os.Exit(1)
		}
		fmt.Printf("User: %v\n", user.Name)
		printFeed(feed)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil

}

func printFeed(feed database.Feed) {
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("URL: %v\n", feed.Url)
	fmt.Printf("Created At: %v\n", feed.CreatedAt)
	fmt.Printf("Updated At: %v\n", feed.UpdatedAt)
	fmt.Println()
}
