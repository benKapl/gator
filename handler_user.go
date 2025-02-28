package main

import (
	"context"
	"fmt"
	"time"

	"github.com/benKapl/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func hanglerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	userParams := database.CreateUserParams{
		ID:        int32(uuid.New()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	s.cfg.CurrentUserName = user.Name
	fmt.Println("User was successfully created!")
	fmt.Printf("%+v", user)

	return nil

}
