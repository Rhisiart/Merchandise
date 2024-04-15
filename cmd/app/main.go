package main

import (
	"log"

	"github.com/Rhisiart/Merchandise/internal/db"
	tables "github.com/Rhisiart/Merchandise/internal/db/tables"
)

func main() {
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	database.Init()

	/*d := &design.Design{
		Name:        "Design 3",
		Description: "Design number 3",
	}

	if err := database.Create(d); err != nil {
		log.Fatal(err)
	}

	log.Printf("Design id = %d", d.DesignId)

	query := &design.Design{
		DesignId: 3,
	}

	if err := database.Read(query); err != nil {
		log.Fatal(err)
	}

	log.Println("description = ", query.Description)
	log.Println("name = ", query.Name)
	log.Println("id = ", query.DesignId)

	c := &customer.Customer{
		Name:    "Customer 1",
		Email:   "customerone@test.com",
		Address: "Viseu",
	}

	if err := database.Create(c); err != nil {
		log.Fatal(err)
	}

	log.Printf("Customer id = %d", c.CustomerId)*/

	customer := &tables.Customer{
		CustomerId: 1,
	}

	if err := database.Read(customer); err != nil {
		log.Fatal(err)
	}

	log.Println("id = ", customer.CustomerId)
	log.Println("name = ", customer.Name)
	log.Println("email = ", customer.Email)
	log.Println("address = ", customer.Address)
}
