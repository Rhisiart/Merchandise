package db

import (
	"context"
	"database/sql"
)

type Order struct {
	OrderId    int64
	CustomerId int64
	CreateOn   string
}

func (order *Order) Create(ctx context.Context, db *sql.DB) error {
	query := `INSERT INTO order (customer_id, createOn)
		VALUES ($1, $2)
		RETURNING order_id`

	err := db.QueryRowContext(
		ctx,
		query,
		order.CustomerId,
		order.CreateOn).Scan(&order.OrderId)

	if err != nil {
		return err
	}

	return nil
}

func (order *Order) Read(ctx context.Context, db *sql.DB) error {
	query := `SELECT customer_id, createOn 
		FROM customer
		WHERE customer_id = $1`

	err := db.QueryRowContext(
		ctx,
		query,
		order.OrderId).Scan(
		&order.CustomerId,
		&order.CreateOn)

	if err != nil {
		return err
	}

	return nil
}

func (order *Order) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (order *Order) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
