package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	// for _, item := range feed.Channel.Item {
	// fmt.Printf("* Item: %v\n", item.Title)
	//
	// }
	fmt.Printf("%+v\n", feed.Channel)
	// fmt.Printf("* Item: %v\n", feed.Channel.Item)
	return nil
}
