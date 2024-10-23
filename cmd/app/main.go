package main

import (
	"github.com/pkstpm/Softdev-Backend/internal/config"
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)
	db.Migrate()
	srv := server.NewEchoServer(cfg, db)

	srv.Start()
}
