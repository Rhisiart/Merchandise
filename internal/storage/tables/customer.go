package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Rhisiart/Merchandise/enum"
	"github.com/Rhisiart/Merchandise/internal/storage"
	"github.com/Rhisiart/Merchandise/types"
)

type Customer struct {
	CustomerId int    `json:"CustomerId,omitempty"`
	Name       string `json:"Name,omitempty"`
	Email      string `json:"Email,omitempty"`
	Address    string `json:"Address,omitempty"`
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
	query := storage.NewQuery(enum.Update, "Customer", customer)
	expression := query.Expression()
	args, e := query.Values()

	if e != nil {
		return e
	}

	q := fmt.Sprintf(`UPDATE customer 
						SET %s
						WHERE customer_id = $1`,
		expression)

	fmt.Printf("query = %s \n", q)

	_, err := db.ExecContext(
		ctx,
		q,
		args...)

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
