package controllers

import (
	// "log"
	// "os"
	"fmt"
	"math/rand"
	"strconv"

	"strings"
	"time"

	"github.com/chukwuka-emi/helpers"
	"github.com/chukwuka-emi/models"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func OpenAccount(c *fiber.Ctx) error {

	account := new(models.UserAccount)

	err := c.BodyParser(account)
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request. Please review your inputs",
			"hint":    err.Error(),
		})
	}

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	name := fmt.Sprintf("%v", claims["name"])
	userId := fmt.Sprintf("%v", claims["userId"])
	email := fmt.Sprintf("%v", claims["email"])

	//generate account number
	rand.Seed(time.Now().Unix())
	account.AccountNumber = strconv.Itoa(rand.Intn(5000000000))

	account.BankNumber = "091"
	account.AvailableBalance = account.AccountBalance - 2000.0
	account.Overdraft = 0.0
	account.GivenName = strings.Split(name, " ")[0]
	account.FamilyName = strings.Split(name, " ")[1]
	account.EmailAddress = email
	account.UserID = userId

	//validate some inputs
	if account.AccountBalance == 0.0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "No initial deposit added",
		})
	}

	if len(account.ContactNumber) < 10 || !strings.HasPrefix(account.ContactNumber, "+") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid phone number",
		})
	}

	if len(account.AddressLine) == 0 || len(account.DateOfBirth) == 0 || len(account.Photo) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "AddressLine, DateOfBirth, Photo cannot be empty",
		})
	}

	if account.AccountType != "savings" && account.AccountType != "current" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Choose a valid account type [savings or current]",
		})
	}

	isValidNIN := helpers.CheckNIN(account.NIN)
	if !isValidNIN {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid NIN",
		})
	}

	//check if account already exist
	existingAccount := models.DB.Where(&models.UserAccount{AccountNumber: account.AccountNumber}).First(new(models.UserAccount))
	if existingAccount.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Account already exist",
		})
	}

	if err := models.DB.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong. Please try again later",
			"hint":    err.Error(),
		})
	}

	helpers.SendAccountOpenMail(account.EmailAddress, account.GivenName, account.AccountNumber, account.AccountType)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":        true,
		"message":        "Congratulations! A mail with your account number has been sent to your account email address",
		"accountDetails": account,
	})

}

func FindAccount(c *fiber.Ctx) error {
	var account models.UserAccount

	err := models.DB.Where("id=?", c.Params("id")).First(&account).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Account not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}

func CloseAccount(c *fiber.Ctx) error {
	var account models.UserAccount

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])

	err := models.DB.Where("account_number=?", c.Params("account_number")).First(&account).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "You don't have the required permission to close an account. Please contact the management",
		})
	}

	e := models.DB.Delete(&account).Error
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured while closing account. Please retry.",
		})
	}

	helpers.SendAccountCloseMail(account.EmailAddress, account.GivenName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Account successfully closed",
	})
}

func ReactivateAccount(c *fiber.Ctx) error {
	var account models.UserAccount

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])

	err := models.DB.Unscoped().Where("account_number=?", c.Params("account_number")).First(&account).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "You don't have the required permission to re-activate this account. Please contact the management",
		})
	}

	models.DB.Unscoped().Delete(&account)

	if err := models.DB.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong. Please try again later",
			"hint":    err.Error(),
		})
	}

	helpers.SendAccountReActivateMail(account.EmailAddress, account.GivenName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Account successfully re-activated",
	})
}
