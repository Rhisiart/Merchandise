package main

import (
	"log"

	"github.com/Rhisiart/Merchandise/internal/db"
	design "github.com/Rhisiart/Merchandise/internal/db/operations"
)

func main() {
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	database.Init()

	design := &design.Design{
		Name:        "Design 2",
		Description: "Design number 2",
	}

	id, err := database.Create(design)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Design id = %d", id)
}
