package main

import (
	"context"
	"fmt"
	"time"

	"github.com/benKapl/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user %w", err)
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFeedsFollowed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user %w", err)
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get followed feeds for %s: %w", user.Name, err)
	}

	if len(followedFeeds) == 0 {
		fmt.Println("No followed feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds for %s:\n", len(followedFeeds), user.Name)
	for _, feed := range followedFeeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
