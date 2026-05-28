package model

import "time"

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string
	Price       float64 `json:"price"`
	ImageURL    string
	CreatedAt   time.Time
}
