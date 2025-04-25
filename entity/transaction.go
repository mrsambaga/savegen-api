package entity

import (
	"time"
)

type Transaction struct {
	ID          			int      			`json:"id" gorm:"primaryKey"`
	UserID      			int       			`json:"user_id" gorm:"foreignKey:UserID"`
	Amount      			float64   			`json:"amount"`
	Detail      			string    			`json:"detail"`
	Date      				time.Time   		`json:"date"`
	CreatedAt   			time.Time 			`json:"created_at"`
	TransactionTypeID  		int                	`json:"transaction_type_id" gorm:"column:type"`
	TransactionCategoryID 	int             	`json:"transaction_category_id" gorm:"column:category"`
	TransactionType    		TransactionType    	`json:"transaction_type" gorm:"foreignKey:TransactionTypeID;references:ID"`
	TransactionCategory 	TransactionCategory `json:"transaction_category" gorm:"foreignKey:TransactionCategoryID;references:ID"`
}