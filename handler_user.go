package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/benKapl/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name) // Check if user is alreay registered
	if err == nil {
		fmt.Fprintf(os.Stderr, "user %s already exists\n", name)
		os.Exit(1)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	s.cfg.CurrentUserName = user.Name
	fmt.Println("User was successfully created!")
	fmt.Printf("%+v\n", user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name) // Checks that user exists
	if err == sql.ErrNoRows {
		fmt.Fprintf(os.Stderr, "user %s does not exist\n", name)
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
