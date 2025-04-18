package entity

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	Detail    string    `json:"detail"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}