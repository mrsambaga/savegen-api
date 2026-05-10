package entity

import "time"

type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	MonthlyBudget *float64  `json:"monthly_budget"`
	CreatedAt     time.Time `json:"created_at"`
}
