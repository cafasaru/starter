package service

import (
	"context"
	"fmt"
	"net/http"
)

// Ping will make a request to the data
func (s *Service) Ping(ctx context.Context) (Health, error) {
	fmt.Println("responding to a ping request")
	return Health{
		Name:    s.Name,
		Version: s.Version,
		Status:  http.StatusOK,
		Message: "ok",
	}, nil
}

// Health pings the database and runs a simple select statment to verify the database is ready for query exec.
func (s *Service) Health(ctx context.Context) (Health, error) {
	// Call the repository ping to ping the database
	fmt.Println("pinging database")

	if err := s.Store.PingDB(ctx); err != nil {
		fmt.Println(err)
		return Health{
			Name:    s.Name,
			Version: s.Version,
			Status:  http.StatusNotAcceptable,
			Message: "failed to ping db",
		}, ErrPingDB
	}

	return Health{
		Name:    s.Name,
		Version: s.Version,
		Status:  http.StatusOK,
		Message: "database ok",
	}, nil
}
