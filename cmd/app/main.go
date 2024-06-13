package main

import (
	"context"
	"log"

	"github.com/Rhisiart/Merchandise/internal/api"
	"github.com/Rhisiart/Merchandise/internal/config"
	"github.com/Rhisiart/Merchandise/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	database := storage.NewDatabase(cfg.Database.DatabaseUrl)
	connErro := database.Connect()

	if connErro != nil {
		log.Fatal(connErro)
	}

	server := api.NewServer(cfg.HTTPServer, database)
	server.Start(ctx)

	defer database.Close()
}
