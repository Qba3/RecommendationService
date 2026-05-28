package repository

import (
	"database/sql"

	"RecommendationService/internal/model"

	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	query := `
		SELECT id, name, description, price, image_url, created_at
		FROM products
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ImageURL,
			&product.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := `
		SELECT id, name, description, price, image_url, created_at
		FROM products
		WHERE id = $1
	`

	var product model.Product

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ImageURL,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetProductsByIDs(
	ids []int,
) ([]model.Product, error) {

	if len(ids) == 0 {
		return []model.Product{}, nil
	}

	query := `
		SELECT id, name, price
		FROM products
		WHERE id = ANY($1)
	`

	rows, err := r.db.Query(
		query,
		pq.Array(ids),
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetPopularProducts() (
	[]model.Product,
	error,
) {

	query := `
		SELECT
			p.id,
			p.name,
			p.price,
			COUNT(oi.product_id) as total
		FROM products p
		LEFT JOIN order_items oi
			ON oi.product_id = p.id
		GROUP BY p.id
		ORDER BY total DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product
		var total int

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&total,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
