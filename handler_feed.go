package main

import (
	"context"
	"fmt"
	"time"

	"github.com/benKapl/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil

}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * CreatedAt:    %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:    %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:    %v\n", feed.Url)
	fmt.Printf(" * UserID:    %v\n", feed.UserID)
}
