package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourname/transport/vehicle/configs"
	"github.com/yourname/transport/vehicle/internal/adapters/grpc/vehiclepb"
)

// Run initializes and starts the HTTP server based on the provided configuration.
// It returns an error if the server fails to start.
func Run(cfg configs.ServerConfig, port vehiclepb.VehicleServiceServer) error {
	log.Printf("Starting server on port %d", cfg.Port)

	router := gin.Default()
	// Add health endpoint
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
	router.GET("vehicles/:id", func(c *gin.Context) {
		ir := vehiclepb.InfoRequest{
			VehicleId: c.Param("id"),
		}
		port.GetVehicleInfo(c, &ir)
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSec) * time.Second,
	}

	log.Println("Starting server on port", cfg.Port)
	// The ListenAndServe call is blocking. It will only return on an
	// unrecoverable error. We return that error to the caller (main).
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %w", server.Addr, err)
	}

	log.Println("Server exited gracefully.")
	return nil
}
