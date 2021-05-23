package models

import (
	"gorm.io/gorm"
)

type Transactions struct{
	gorm.Model
	TransactionType string `json:"transaction_type"`
	TransactionTitle string `json:"transaction_title"`
	Amount string `json:"amount"`
	Status string `json:"status"`
	SourceAccount string `json:"source_account"`
	UserID uint
}