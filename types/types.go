package types

import "database/sql"

type IOperation interface {
	Create(db *sql.DB) (int64, error)
	Read(id int64) error
	Update() error
	Delete(id int64) error
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
