package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserCreateResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
