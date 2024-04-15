package db

import "database/sql"

type OrderItem struct {
	OrderItemId int64
	OrderId     int64
	ClothingId  string
	Quantity    int
}

func (orderItem *OrderItem) Create(db *sql.DB) error {
	query := `INSERT INTO orderitem (order_id, clothing_id, quantity)
		VALUES ($1, $2, $3)
		RETURNING order_item_id`

	err := db.QueryRow(
		query,
		orderItem.OrderId,
		orderItem.ClothingId,
		orderItem.Quantity).Scan(&orderItem.OrderItemId)

	if err != nil {
		return err
	}

	return nil
}

func (orderItem *OrderItem) Read(db *sql.DB) error {
	query := `SELECT order_id, clothing_id, quantity 
		FROM orderitem
		WHERE order_item_id = $1`

	err := db.QueryRow(
		query,
		orderItem.OrderItemId).Scan(
		&orderItem.OrderId,
		&orderItem.ClothingId,
		&orderItem.Quantity)

	if err != nil {
		return err
	}

	return nil
}

func (orderItem *OrderItem) Update(db *sql.DB) error {
	return nil
}

func (orderItem *OrderItem) Delete(db *sql.DB) error {
	return nil
}
