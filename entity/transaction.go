package entity

import (
	"time"
)

type Transaction struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"foreignKey:UserID"`
	Amount      float64   `json:"amount"`
	Type        int       `json:"type" gorm:"foreignKey:TransactionType"`
	Category    int       `json:"category" gorm:"foreignKey:TransactionCategory"`
	Detail      string    `json:"detail"`
	CreatedAt   time.Time `json:"created_at"`
}