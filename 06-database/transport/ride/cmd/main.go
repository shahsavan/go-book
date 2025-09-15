package main

import (
	"log"

	"github.com/yourname/transport/ride/configs"
	"github.com/yourname/transport/ride/internal/adapters/httpserver"
	"github.com/yourname/transport/ride/internal/adapters/repository"
	"github.com/yourname/transport/ride/internal/core/ports"
)

func main() {
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Mask secrets before logging
	safe := *cfg
	safe.Database.Password = "<redacted>"
	log.Printf("Loaded config: %+v", safe)
	var repo ports.AssignmentRepository
	repo = repository.NewSQLAssignmentRepository(cfg.Database.OpenDB())
	httpserver.Run(*cfg)
}
