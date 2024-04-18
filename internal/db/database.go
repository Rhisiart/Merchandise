package db

import (
	"database/sql"

	"github.com/Rhisiart/Merchandise/types"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type Database struct {
	url      string
	database *sql.DB
}

func NewDatabase(databaseUrl string) *Database {
	return &Database{
		url: databaseUrl,
	}
}

func (db *Database) Connect() error {
	database, err := sql.Open(driverName, db.url)

	if err != nil {
		return err
	}

	db.database = database
	return nil
}

func (db *Database) Create(operation types.Operation) error {
	return operation.Create(db.database)
}

func (db *Database) Read(operation types.Operation) error {
	return operation.Read(db.database)
}

func (db *Database) Update(operation types.Operation) error {
	return nil
}

func (db *Database) Delete(operation types.Operation) error {
	return nil
}

func (db *Database) Close() error {
	return db.database.Close()
}
