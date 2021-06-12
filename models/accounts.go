package models

import (
	"gorm.io/gorm"
)

type UserAccount struct {
	gorm.Model
	AccountNumber    string  `json:"account_number"`
	BankNumber       string  `json:"bank_number"`
	GivenName        string  `json:"given_name"`
	FamilyName       string  `json:"family_name"`
	EmailAddress     string  `json:"email_address"`
	DateOfBirth      string  `json:"date_of_birth"`
	Photo            string  `json:"photo"`
	NIN              string  `json:"nin"`
	ContactNumber    string  `json:"contact_number"`
	AddressLine      string  `json:"addressline"`
	AccountBalance   float64 `json:"account_balance"`
	Overdraft        float64 `json:"overdraft"`
	AvailableBalance float64 `json:"available_balance"`
	AccountType      string  `json:"account_type"`
	UserID           string  `json:"user_id"`
}
