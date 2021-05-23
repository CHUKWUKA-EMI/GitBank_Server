package models

import (
	"gorm.io/gorm"
)

type UserAccount struct{
	gorm.Model
	AccountNumber string `json:"account_number"`
	BankNumber string `json:"bank_number"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	DateOfBirth string `json:"date_of_birth"`
	Photo string `json:"photo"`
  NIN string `json:"nin"`
  ContactNumber1 string `json:"contact_number1"`
  ContactNumber2 string `json:"contact_number2"`
  EmailAddress string `json:"email_address"`
  AddressLine1 string `json:"addressline1"`
  AddressLine2 string `json:"addressline2"`
  AddressLine3 string `json:"addressline3"`
  PostalCode string `json:"postalcode"`
	AccountBalance float64 `json:"account_balance"`
	Overdraft float64 `json:"overdraft"`
	AvailableBalance float64 `json:"available_balance"`
	UserID uint
}
