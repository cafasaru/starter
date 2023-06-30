package db

import "context"

// PingDB pings the database and return an error on failure
func (d *Database) PingDB(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
