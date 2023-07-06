// Package service is responsible for providing the status of our microservice
package service

import (
	"context"
	"errors"
)

var (
	ErrPingDB = errors.New("failed to ping the database")
)

// Health is a representation of the health struct which will provide service status
type Health struct {
	Name    string `json:"service"`
	Version string `json:"version"`
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

// Store represents the interface used to interact with the repository
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
func NewService(name, version string, store Store) *Service {
	return &Service{
		Name:    name,
		Version: version,
		Store:   store,
	}
}
