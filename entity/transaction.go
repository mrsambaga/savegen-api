package entity

import (
	"time"
)

type Transaction struct {
	ID          		int      `json:"id" gorm:"primaryKey"`
	UserID      		int       `json:"user_id" gorm:"foreignKey:UserID"`
	Amount      		float64   `json:"amount"`
	Detail      		string    `json:"detail"`
	CreatedAt   		time.Time `json:"created_at"`
	TransactionType 	TransactionType `json:"transaction_type" gorm:"foreignKey:TransactionType"`
	TransactionCategory TransactionCategory `json:"transaction_category" gorm:"foreignKey:TransactionCategory"`
}