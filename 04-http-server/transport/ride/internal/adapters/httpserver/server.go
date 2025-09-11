package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yourname/transport/ride/configs"
)

// Run initializes and starts the HTTP server based on the provided configuration.
// It returns an error if the server fails to start.
func Run(cfg configs.Config) error {
	log.Printf("Starting server on port %d", cfg.Server.Port)

	// Create a simple handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the ride service!")
	})

	// Liveness probe endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Ok")
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeoutSec) * time.Second,
	}

	// The ListenAndServe call is blocking. It will only return on an
	// unrecoverable error. We return that error to the caller (main).
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %w", server.Addr, err)
	}

	log.Println("Server exited gracefully.")
	return nil
}
