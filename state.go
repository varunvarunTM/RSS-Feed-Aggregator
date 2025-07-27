package main

import (
	"RSS-feed-aggregator/config"
	"RSS-feed-aggregator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

