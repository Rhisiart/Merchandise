package db

import (
	"context"
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

func (db *Database) Create(ctx context.Context, operation types.Table) error {
	return operation.Create(ctx, db.database)
}

func (db *Database) Read(ctx context.Context, operation types.Table) error {
	return operation.Read(ctx, db.database)
}

func (db *Database) ReadAll(ctx context.Context, operation types.Table, list *[]types.Table) error {
	return operation.ReadAll(ctx, db.database, list)
}

func (db *Database) Update(ctx context.Context, operation types.Table) error {
	return operation.Update(ctx, db.database)
}

func (db *Database) Delete(ctx context.Context, operation types.Table) error {
	return operation.Delete(ctx, db.database)
}

func (db *Database) Close() error {
	return db.database.Close()
}
