package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Rhisiart/Merchandise/types"
)

type Customer struct {
	CustomerId int    `json:"CustomerId"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Address    string `json:"Address"`
}

func (customer *Customer) Create(ctx context.Context, db *sql.DB) error {
	query := `INSERT INTO customer (name, email, address)
		VALUES ($1, $2, $3)
		RETURNING customer_id`

	err := db.QueryRowContext(
		ctx,
		query,
		customer.Name,
		customer.Email,
		customer.Address).Scan(&customer.CustomerId)

	if err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Read(ctx context.Context, db *sql.DB) error {
	query := `SELECT name, email, address 
		FROM customer
		WHERE customer_id = $1`

	err := db.QueryRowContext(
		ctx,
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

func (customer *Customer) ReadAll(ctx context.Context, db *sql.DB, list *[]types.Table) error {
	query := `SELECT customer_id, name, email, address
			FROM customer`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		newCustomer := new(Customer)

		if err := rows.Scan(
			&newCustomer.CustomerId,
			&newCustomer.Name,
			&newCustomer.Email,
			&newCustomer.Address); err != nil {
			return err
		}

		*list = append(*list, newCustomer)
	}

	return nil
}

func (customer *Customer) Update(ctx context.Context, db *sql.DB) error {
	columns := "name = $1, email = $2, address = $3"

	query := fmt.Sprintf(`UPDATE customer 
						SET %s
						WHERE customer_id = $4`,
		columns)

	_, err := db.ExecContext(
		ctx,
		query,
		customer.Name,
		customer.Email,
		customer.Address,
		customer.CustomerId)

	if err != nil {
		return err
	}

	return nil
}

func (customer *Customer) Delete(ctx context.Context, db *sql.DB) error {
	query := `DELETE
			FROM customer
			WHERE customer_id = $1`

	result, err := db.ExecContext(ctx, query, customer.CustomerId)

	if err != nil {
		return fmt.Errorf("error deleting customer: %w", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("no rows were deleted")
	}

	return nil
}
