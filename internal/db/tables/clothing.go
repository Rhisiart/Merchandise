package db

import (
	"database/sql"
)

type Clothing struct {
	ClothingId int64
	DesignId   int64
	Name       string
	Price      int
}

func (clothing *Clothing) Create(db *sql.DB) error {
	query := `INSERT INTO clothing (design_id, name, price)
		VALUES ($1, $2, $3)
		RETURNING clothing_id`

	err := db.QueryRow(
		query,
		clothing.DesignId,
		clothing.Name,
		clothing.Price).Scan(&clothing.ClothingId)

	if err != nil {
		return err
	}

	return nil
}

func (clothing *Clothing) Read(db *sql.DB) error {
	query := `SELECT design_id, name, price
		FROM clothing
		WHERE clothing_id = $1`

	err := db.QueryRow(
		query,
		clothing.ClothingId).Scan(
		&clothing.DesignId,
		&clothing.Name,
		&clothing.Price)

	if err != nil {
		return err
	}

	return nil
}

func (clothing *Clothing) Update(db *sql.DB) error {
	return nil
}

func (clothing *Clothing) Delete(db *sql.DB) error {
	return nil
}
