package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cafasaru/starter/internal/health"
	"github.com/gorilla/mux"
)

// Handler is the struct we use for interacting with the transport layer of our app
// when a NewHandler is constructed it applies the routes found in the mapRoutes method
type Handler struct {
	Router  *mux.Router
	Service Service
	Server  *http.Server
}

// Service represents the interface which is the logical layer representation of the microservice
// Add the service methods here
type Service interface {
	Health(context.Context) (health.Health, error)
}

// NewHandler is how a new handler with its routes is constructed.
func NewHandler(service Service) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

// mapRoutes is used to map the route to the handler used for the request
func (h *Handler) mapRoutes() {

	h.Router.HandleFunc("/api/v1/liveness", h.Health).Methods("GET")
}

// Serve is responsible for starting up the server and gracefully shutting down
// if a shutdown signal is received
func (h *Handler) Serve() error {

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")

	return nil
}
