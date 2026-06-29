package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect initializes and verifies the connection pool with explicit pool settings.
func Connect(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	// 1. Parse the connection string into a config struct(provided by ParseConfig)
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("invalid database configuration URL: %w", err)
	}

	// 2. Adjust the pool settings directly on the struct
	// Go uses time.Duration for lifetimes, making it explicit and readable.
	config.MaxConns = 25                      // Maximum number of connections in the pool
	config.MinConns = 5                       // Minimum number of idle connections to maintain
	config.MaxConnIdleTime = 15 * time.Minute // How long an idle connection lives before being closed
	config.MaxConnLifetime = 1 * time.Hour    // Maximum lifetime of any connection

	// 3. Create the pool using the modified config
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// 4. Validates the actual connection to the server
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // Clean up the pool if the ping fails
		return nil, fmt.Errorf("database network connection failed: %w", err)
	}

	return pool, nil
}