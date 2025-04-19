package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
}

type UserCreateResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
