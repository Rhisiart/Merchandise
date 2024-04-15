package types

import (
	"database/sql"
)

type Operation interface {
	Create(db *sql.DB) error
	Read(db *sql.DB) error
	Update(db *sql.DB) error
	Delete(db *sql.DB) error
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
