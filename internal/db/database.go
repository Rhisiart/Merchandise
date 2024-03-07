package database

import (
	"database/sql"
	"fmt"
)

const (
	driverName = "postgres"
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "bd"
	dbname     = "Merchandise"
)

type database struct {
	db *sql.DB
}

func NewDatabase() (*database, error) {
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

	return &database{
		db: db,
	}, nil
}
