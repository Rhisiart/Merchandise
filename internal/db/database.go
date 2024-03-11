package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Rhisiart/Merchandise/types"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "bd"
	dbname     = "Merchandise"
)

type Database struct {
	database *sql.DB
}

func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open(driverName, connString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		database: db,
	}, nil
}

func (db *Database) Init() {
	log.Printf("Database initialized...")
	status := db.database.Stats()
	log.Print("Database status = ", status)
}

func (db *Database) Create(operation types.IOperation) (int64, error) {
	return operation.Create(db.database)
}

func (db *Database) Close() {
	db.database.Close()
	log.Printf("Databse closed")
}
