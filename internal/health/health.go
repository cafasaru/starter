// Package health is responsible for providing the status of our microservice health
// also is used for k8s status
package health

import (
	"context"
	"fmt"
	"net/http"
)

// Health is a representation of a health struct which will provide db status
type Health struct {
	Name    string `json:"service"`
	Version string `json:"version"`
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

// Store represents the interface with which we interact with the repository layer
type Store interface {
	// PingDB pings the database
	PingDB(ctx context.Context) error
}

// Service is the struct used to interact with the service
type Service struct {
	Name    string
	Version string
	Store   Store
}

// NewService is how a new service is constructred
func NewService(n, v string, s Store) *Service {
	return &Service{
		Name:    n,
		Version: v,
		Store:   s,
	}
}

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
	fmt.Println("call the database for status")

	if err := s.Store.PingDB(ctx); err != nil {
		fmt.Println(err)
		return Health{
			Name:    s.Name,
			Version: s.Version,
			Status:  http.StatusNotAcceptable,
			Message: "failed to ping db",
		}, nil
	}

	return Health{
		Name:    s.Name,
		Version: s.Version,
		Status:  http.StatusOK,
		Message: "database ok",
	}, nil
}
