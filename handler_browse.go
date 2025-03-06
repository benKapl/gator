package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/benKapl/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var err error

	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("<limit> must be a number: %w", err)
		}
	} else {
		limit = 2
	}

	posts, err := s.db.GetPostsForUser(context.TODO(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Couldn't retrieve posts: %w", err)
	}

	fmt.Printf("Found %d posts for %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("* %s, %s\n", post.Title, post.Url)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}
