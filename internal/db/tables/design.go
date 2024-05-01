package db

import (
	"context"
	"database/sql"
)

type Design struct {
	DesignId    int
	Name        string
	Description string
}

func (design *Design) Create(ctx context.Context, db *sql.DB) error {
	query := `INSERT INTO design (name, description)
		VALUES ($1, $2)
		RETURNING design_id`

	err := db.QueryRowContext(
		ctx,
		query,
		design.Name,
		design.Description).Scan(&design.DesignId)

	if err != nil {
		return err
	}

	return nil
}

func (design *Design) Read(ctx context.Context, db *sql.DB) error {
	query := `SELECT name, description 
		FROM design
		WHERE design_id = $1`

	err := db.QueryRowContext(
		ctx,
		query,
		design.DesignId).Scan(&design.Name, &design.Description)

	if err != nil {
		return err
	}

	return nil
}

func (design *Design) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (design *Design) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
