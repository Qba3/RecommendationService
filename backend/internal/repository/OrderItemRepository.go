package repository

import (
	"database/sql"

	"RecommendationService/internal/model"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) CreateOrderItem(orderID int, productID int, quantity int) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, orderID, productID, quantity)
	return err
}

func (r *OrderItemRepository) GetItemsByOrderID(orderID int) ([]model.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity
		FROM order_items
		WHERE order_id = $1
	`

	rows, err := r.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []model.OrderItem

	for rows.Next() {
		var item model.OrderItem

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
