package db

import (
	"database/sql"
)

type Design struct {
	DesignId    int64
	Name        string
	Description string
}

func (design *Design) Create(db *sql.DB) error {
	query := `INSERT INTO design (name, description)
		VALUES ($1, $2)
		RETURNING design_id`

	err := db.QueryRow(
		query,
		design.Name,
		design.Description).Scan(&design.DesignId)

	if err != nil {
		return err
	}

	return nil
}

func (design *Design) Read(db *sql.DB) error {
	query := `SELECT name, description 
		FROM design
		WHERE design_id = $1`

	err := db.QueryRow(
		query,
		design.DesignId).Scan(&design.Name, &design.Description)

	if err != nil {
		return err
	}

	return nil
}

func (design *Design) Update(db *sql.DB) error {
	return nil
}

func (design *Design) Delete(db *sql.DB) error {
	return nil
}
