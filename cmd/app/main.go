package main

import (
	"log"

	database "github.com/Rhisiart/Merchandise/internal/db"
)

func main() {
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	db.Init()

}
