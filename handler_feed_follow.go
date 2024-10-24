package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hawkaii/blogator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	user, err := s.dbQueries.GetUser(context.Background(), s.cfg.Current_User_Name)

	if err != nil {
		return err
	}

	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %v <name>", cmd.Name)
		os.Exit(1)
	}

	url := cmd.Args[0]

	feed, err := s.dbQueries.GetFeedByURL(context.Background(), url)

	if err != nil {
		fmt.Print("Feed not found %w", err)
		os.Exit(1)
	}

	ff_row, err := s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		fmt.Print("Feed follow not created %w", err)
		os.Exit(1)
	}

	fmt.Println("Followed: details below")
	printFeedFollow(ff_row.UserName, ff_row.FeedName)
	fmt.Println("=====================================")

	return nil
}

func handlerFollowing(s *state, _ command) error {
	user, err := s.dbQueries.GetUser(context.Background(), s.cfg.Current_User_Name)

	if err != nil {
		return err
	}

	feed_follows, err := s.dbQueries.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		fmt.Print("Feed follow not found %w", err)
		os.Exit(1)
	}

	fmt.Println("Following list below:")
	for _, feed_follow := range feed_follows {
		fmt.Printf("* %s\n", feed_follow.FeedName)
		fmt.Println("=====================================")
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %v <name>", cmd.Name)
		os.Exit(1)
	}

	url := cmd.Args[0]

	feed, err := s.dbQueries.GetFeedByURL(context.Background(), url)

	if err != nil {
		fmt.Print("Feed not found %w", err)
		os.Exit(1)
	}

	_, err = s.dbQueries.DeleteFeedFollowByUserAndUrl(context.Background(),
		database.DeleteFeedFollowByUserAndUrlParams{
			UserID: user.ID,
			Url:    url,
		})

	if err != nil {
		fmt.Print("Feed follow not deleted %w", err)
		os.Exit(1)
	}

	fmt.Println("Unfollowed: details below")
	printFeedFollow(user.Name, feed.Name)
	fmt.Println("=====================================")

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("User: %s\n", username)
	fmt.Printf("Feed: %s\n", feedname)
}
