package main

import (
	"log"

	"github.com/Rhisiart/Merchandise/internal/api"
	"github.com/Rhisiart/Merchandise/internal/config"
	"github.com/Rhisiart/Merchandise/internal/db"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	database := db.NewDatabase(cfg.Database.DatabaseUrl)

	connErro := database.Connect()

	if connErro != nil {
		log.Fatal(connErro)
	}

	server := api.NewServer(cfg.HTTPServer, database)

	server.Start()
}
