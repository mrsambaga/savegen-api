package entity

import "time"

type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PasswordHash  *string   `json:"-"`
	IsGuest       bool      `json:"is_guest"`
	GoogleSub     *string   `json:"-"`
	MonthlyBudget *float64  `json:"monthly_budget"`
	CreatedAt     time.Time `json:"created_at"`
}
