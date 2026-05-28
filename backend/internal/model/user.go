package model

import "time"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt time.Time
}
