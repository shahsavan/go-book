package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourname/transport/vehicle/configs"
	"github.com/yourname/transport/vehicle/internal/adapters/grpc/vehiclepb"
	"github.com/yourname/transport/vehicle/internal/adapters/httpserver"
)

func main() {
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	safe := *cfg
	safe.Database.Password = "<redacted>"
	log.Printf("Loaded config: %+v", safe)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	var vs vehiclepb.VehicleServiceServer
	httpserver.Run(cfg.Server, vs)
}
