package db

import (
	"context"
	"database/sql"
)

type Clothing struct {
	ClothingId int64
	DesignId   int64
	Name       string
	Price      int
}

func (clothing *Clothing) Create(ctx context.Context, db *sql.DB) error {
	query := `INSERT INTO clothing (design_id, name, price)
		VALUES ($1, $2, $3)
		RETURNING clothing_id`

	err := db.QueryRowContext(
		ctx,
		query,
		clothing.DesignId,
		clothing.Name,
		clothing.Price).Scan(&clothing.ClothingId)

	if err != nil {
		return err
	}

	return nil
}

func (clothing *Clothing) Read(ctx context.Context, db *sql.DB) error {
	query := `SELECT design_id, name, price
		FROM clothing
		WHERE clothing_id = $1`

	err := db.QueryRowContext(
		ctx,
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

func (clothing *Clothing) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (clothing *Clothing) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
