package types

import (
	"context"
	"database/sql"
)

type Operation interface {
	Create(ctx context.Context, db *sql.DB) error
	Read(ctx context.Context, db *sql.DB) error
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
