package helpers

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/chukwuka-emi/models"
)

type UserErrors struct {
	Err       bool   `json:"error"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

// IsEmpty checks if a string is empty
func IsEmpty(str string) (bool, string) {
	if valid.HasWhitespaceOnly(str) && str != "" {
		return true, "Must not be empty"
	}

	return false, ""
}

//CheckNIN checks if a string is a valid NIN
func CheckNIN(str string) bool {
	if !valid.IsInt(str) || len(str) < 10 || !valid.IsNumeric(str) {
		return false
	}

	return true
}

// ValidateRegister func validates the body of user for registration
func ValidateRegister(u *models.User) UserErrors {
	e := UserErrors{}
	e.Err, e.FirstName = IsEmpty(u.FirstName)
	e.Err, e.LastName = IsEmpty(u.LastName)

	if !valid.IsEmail(u.Email) {
		e.Err, e.Email = true, "Must be a valid email"
	}

	if !(len(u.Password) >= 8 && valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password)) {
		e.Err, e.Password = true, "Length of password should be atleast 8 and it must be a combination of uppercase letters and lowercase letters"
	}

	return e
}
