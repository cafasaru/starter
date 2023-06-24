// Package database sets up a default pg db with sqlx
// this package is also responsible for migrating our database
package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database is the struct used to provide a client connection to the database
type Database struct {
	Client *sqlx.DB
}

// NewDatabase is responsible for constructing a database connection and pinging the database
func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %w", err)
	}

	return &Database{
		Client: dbConn,
	}, nil
}

// PingDB pings the database and return an error on failure
func (d *Database) PingDB(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
