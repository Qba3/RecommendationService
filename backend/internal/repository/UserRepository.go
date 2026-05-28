package repository

import (
	"database/sql"

	"RecommendationService/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(username string) (int, error) {
	var userID int

	query := `
		INSERT INTO users (username)
		VALUES ($1)
		RETURNING id
	`

	err := r.db.QueryRow(query, username).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, created_at
		FROM users
		WHERE username = $1
	`

	var user model.User

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	query := `
		SELECT id, username, created_at
		FROM users
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
