package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/benKapl/gator/internal/database"
	"github.com/google/uuid"
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

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeed(s)
	}
}

func scrapeFeed(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Couldn't identify next feed: %v", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't fetch feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		savePost(s, feed.ID, item)
		// fmt.Printf("pub date, %v\n", item.PubDate)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

func savePost(s *state, feedId uuid.UUID, feedItem RSSItem) {
	pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", feedItem.PubDate)
	if err != nil {
		log.Printf("Couldn't parse publication date %v\n", err)
		return
	}

	_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       feedItem.Title,
		Url:         feedItem.Link,
		Description: toNullString(feedItem.Description),
		PublishedAt: sql.NullTime{Time: pubDate, Valid: true},
		FeedID:      feedId,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {

		} else {
			// Log other types of errors
			log.Printf("Error creating post: %v\n", err)
		}
	}
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false} // Null value
	}
	return sql.NullString{String: s, Valid: true}
}
