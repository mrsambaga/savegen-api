package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50,string"`
	Email    string `json:"email" binding:"required,email"`
}

type UserCreateResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
