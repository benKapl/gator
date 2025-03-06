package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return fmt.Errorf("couldn't fetch feed: %w", err)
	// }

	// fmt.Printf("Feed: %+v\n", feed)
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	err = scrapeFeed(s, context.Background())
	if err != nil {
		return fmt.Errorf("couldn't scrape feed: %w", err)
	}

	return nil
}
