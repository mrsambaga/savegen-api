package dto

import "time"

type TransactionRequest struct {
	UserID      	*int    `json:"user_id" form:"user_id"`
	TransactionType *string `json:"transaction_type" form:"transaction_type"`
}

type TransactionCreateRequest struct {
	UserID      		*int    	`json:"user_id" form:"user_id"`
	Amount      		*float64 	`json:"amount" form:"amount"`
	Detail      		*string 	`json:"detail" form:"detail"`
	TransactionType 	*string 	`json:"transaction_type" form:"transaction_type"`
	TransactionCategory *string 	`json:"transaction_category" form:"transaction_category"`
	Date      			*string 	`json:"date" form:"date"`
}

type TransactionResponse struct {
	ID          		int       `json:"id"`
	UserID      		int       `json:"user_id"`
	Amount      		float64   `json:"amount"`
	TransactionType     string    `json:"transaction_type"`
	TransactionCategory string    `json:"transaction_category"`
	Detail      		string    `json:"detail"`
	Date      			string    `json:"date"` 
	CreatedAt   		time.Time `json:"created_at"`
}

type TransactionListResponse struct {
	Code     string              	`json:"code"`
	Messages string              	`json:"messages"`
	Data     []TransactionResponse 	`json:"data"`
} 