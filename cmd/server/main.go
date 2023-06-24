package main

import (
	"fmt"

	"github.com/cafasaru/starter/internal/db"
)

const (
	// Version for the microservice
	Version = "0.0.1"
)

// Run is responsible for setting up and running the microservice
func Run() error {
	fmt.Println("running service")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

// main is the main entrypoint of the microservice
func main() {
	fmt.Println("starting server")

	if err := Run(); err != nil {
		fmt.Println(err.Error())
	}
}
