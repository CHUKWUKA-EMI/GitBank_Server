package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Verified     bool   `json:"verified"`
	Role         string `json:"role"`
	Accounts     []UserAccount
	Transactions []Transactions
}
