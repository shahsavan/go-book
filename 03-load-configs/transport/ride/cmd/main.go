package main

import (
	"log"

	"github.com/yourname/transport/ride/configs"
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
}
