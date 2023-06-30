package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// MigrateDB is used to migrate the database
// it looks in the migrations folder and executes the up files in order
// ex: 0001_first.up.sql > 0002_second.up.sql
func (d *Database) MigrateDB() error {
	fmt.Println("migrating database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to greate the postgres migration driver %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations %w", err)
		}
	}

	fmt.Println("successfully migrated the database")

	return nil
}
