package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourname/transport/ride/configs"
	"github.com/yourname/transport/ride/internal/adapters/httpserver"
	"github.com/yourname/transport/ride/internal/adapters/repository"
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

	assignmentRepo := repository.NewSQLAssignmentRepository(db)

	httpserver.Run(cfg.Server, assignmentRepo)
}
