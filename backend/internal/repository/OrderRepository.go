package repository

import (
	"database/sql"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(userID int) (int, error) {
	var orderID int

	query := `
		INSERT INTO orders (user_id)
		VALUES ($1)
		RETURNING id
	`

	err := r.db.QueryRow(query, userID).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (r *OrderRepository) AddOrderItem(orderID int, productID int, quantity int) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, orderID, productID, quantity)
	return err
}

func (r *OrderRepository) GetUserOrders(userID int) (*sql.Rows, error) {
	query := `
		SELECT id, user_id, created_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	return r.db.Query(query, userID)
}

func (r *OrderRepository) GetOrderItems(orderID int) (*sql.Rows, error) {
	query := `
		SELECT id, order_id, product_id, quantity
		FROM order_items
		WHERE order_id = $1
	`

	return r.db.Query(query, orderID)
}

func (r *OrderRepository) GetUserPurchasedProducts(
	userID int,
) ([]int, error) {

	query := `
		SELECT DISTINCT oi.product_id
		FROM orders o
		JOIN order_items oi
			ON oi.order_id = o.id
		WHERE o.user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var productIDs []int

	for rows.Next() {

		var productID int

		err := rows.Scan(&productID)
		if err != nil {
			return nil, err
		}

		productIDs = append(productIDs, productID)
	}

	return productIDs, nil
}

func (r *OrderRepository) GetAllUserProducts() (
	map[int][]int,
	error,
) {

	query := `
		SELECT o.user_id, oi.product_id
		FROM orders o
		JOIN order_items oi
			ON oi.order_id = o.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := map[int][]int{}

	for rows.Next() {

		var userID int
		var productID int

		err := rows.Scan(
			&userID,
			&productID,
		)
		if err != nil {
			return nil, err
		}

		result[userID] = append(
			result[userID],
			productID,
		)
	}

	return result, nil
}
