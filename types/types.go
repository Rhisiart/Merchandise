package types

import (
	"context"
	"database/sql"
)

type Table interface {
	Create(ctx context.Context, db *sql.DB) error
	Read(ctx context.Context, db *sql.DB) error
	ReadAll(ctx context.Context, db *sql.DB, list *[]Table) error
	Update(ctx context.Context, db *sql.DB) error
	Delete(ctx context.Context, db *sql.DB) error
}

type Controllers interface {
}

type ClothingType struct {
	ClothingTypeId int64
	Name           string
	Description    string
}

type Clothing struct {
	ClothingId int64
	DesignId   int64
	Name       string
	Price      int
}
