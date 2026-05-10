package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateRequest struct {
	Username      *string  `json:"username" binding:"omitempty,min=3,max=50"`
	MonthlyBudget *float64 `json:"monthly_budget" binding:"omitempty,gte=0"`
}

type UserCreateResponse struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	MonthlyBudget *float64  `json:"monthly_budget"`
	CreatedAt     time.Time `json:"created_at"`
}
