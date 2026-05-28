package model

import "time"

type Order struct {
	ID        int
	UserID    int
	CreatedAt time.Time
}
