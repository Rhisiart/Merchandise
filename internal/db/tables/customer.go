package db

import "database/sql"

type Customer struct {
	CustomerId int64
	Name       string
	Email      string
	Address    string
}

func (customer *Customer) Create(db *sql.DB) error {
	query := `INSERT INTO customer (name, email, address)
		VALUES ($1, $2, $3)
		RETURNING customer_id`

	err := db.QueryRow(
		query,
		customer.Name,
		customer.Email,
		customer.Address).Scan(&customer.CustomerId)

	if err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Read(db *sql.DB) error {
	query := `SELECT name, email, address 
		FROM customer
		WHERE customer_id = $1`

	err := db.QueryRow(
		query,
		customer.CustomerId).Scan(
		&customer.Name,
		&customer.Email,
		&customer.Address)

	if err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Update(db *sql.DB) error {
	return nil
}

func (customer *Customer) Delete(db *sql.DB) error {
	return nil
}
