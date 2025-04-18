package dto

import "time"

type TransactionRequest struct {
	UserID      *int    `json:"user_id" form:"user_id"`
	TypeName    *string `json:"type_name" form:"type_name"`
}

type TransactionResponse struct {
	ID          		int       `json:"id"`
	UserID      		int       `json:"user_id"`
	Amount      		float64   `json:"amount"`
	TransactionType     string    `json:"transaction_type"`
	TransactionCategory string    `json:"transaction_category"`
	Detail      		string    `json:"detail"`
	CreatedAt   		time.Time `json:"created_at"`
}

type TransactionListResponse struct {
	Code     string              	`json:"code"`
	Messages string              	`json:"messages"`
	Data     []TransactionResponse 	`json:"data"`
} 