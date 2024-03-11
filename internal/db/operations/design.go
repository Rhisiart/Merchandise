package db

import (
	"database/sql"
)

type Design struct {
	DesignId    int64
	Name        string
	Description string
}

func (design *Design) Create(db *sql.DB) (int64, error) {
	query := `INSERT INTO design (name, description)
	VALUES ($1, $2)
	RETURNING design_id`

	var designId int64
	err := db.QueryRow(
		query,
		design.Name,
		design.Description).Scan(&designId)

	if err != nil {
		return 0, err
	}

	return designId, nil
}

func (design *Design) Read(id int64) error {
	return nil
}

func (design *Design) Update() error {
	return nil
}

func (design *Design) Delete(id int64) error {
	return nil
}
