package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhaatha/go-e-commerce-api/internal/config"
)

func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create database pool: %w", err)
	}

	// Check DB connection
	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %w", err)
	}

	slog.Info("connected to the database")
	return dbPool, nil
}
