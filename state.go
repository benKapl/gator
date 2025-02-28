package main

import (
	"github.com/benKapl/gator/internal/config"
	"github.com/benKapl/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
