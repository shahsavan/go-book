package main

import (
	"flag"
	"log"

	"github.com/yourname/transport/ride/configs"
	"github.com/yourname/transport/ride/internal/adapters/httpserver"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "path to config file")
	flag.Parse()
	cfg, err := configs.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	if err := httpserver.Run(*cfg); err != nil {
		log.Fatalf("application run failed: %v", err)
	}
}
