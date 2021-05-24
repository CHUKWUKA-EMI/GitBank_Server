package controllers

import (
	// "log"
	"os"
	"time"

	"github.com/chukwuka-emi/helpers"
	"github.com/chukwuka-emi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

type Claims struct {
 jwt.StandardClaims
 ID uint `gorm:"primaryKey"`
 }

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(c *fiber.Ctx) error{
  user := new(models.User)

	err := c.BodyParser(user)
	if err != nil{
		return c.JSON(fiber.Map{
			"error": true,
			"message":"Invalid request. Please review your inputs",
		})
	}

	errors := helpers.ValidateRegister(user)
	if errors.Err{
		return c.JSON(errors)
	}

	existingEmail := models.DB.Where(&models.User{Email: user.Email}).First(new(models.User))
	if existingEmail.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message":"Email already exist",
		})
	}

	existingNames := models.DB.Where(&models.User{FirstName: user.FirstName, LastName: user.LastName}).First(new(models.User))
	if existingNames.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message":"Names already taken",
		})
	}

	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password,14)
	if err !=nil{
		return c.JSON(fiber.Map{
			"error":true,
			"message":err.Error(),
		})
	}

	user.Password = string(hashedPassword)
	user.Verified = false

	if err := models.DB.Create(&user).Error;err !=nil{
		return c.JSON(fiber.Map{
			"error":true,
			"message":"Something went wrong. Please try again later",
		})
	}
  
	token, err := helpers.GenerateVerificationToken(user.ID)
	if err !=nil{
		return c.JSON(fiber.Map{
			"error":true,
			"message":err.Error(),
		})
	}

  
	backend_url := os.Getenv("BACKEND_URL")
	url := backend_url+"/api/v1/user/verify/"+token
	helpers.SendMail(user.Email, user.FirstName,url)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"access_token": token,
	})
}

func VerifyEmail(c *fiber.Ctx) error {
	var user models.User
	paramToken := c.Params("token")
	if len(paramToken)==0{
		c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error":true,
			"message":"No token provided",
		})
	}

 	tokenString, err := jwt.ParseWithClaims(paramToken, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return jwtKey, nil
	})

	if err !=nil{
		return  err
	}
	
	claims, ok := tokenString.Claims.(*Claims)
	if !ok && !tokenString.Valid{
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": true,
			"message":"Couldn't parse claims",
		})
	}

	if claims.ExpiresAt < time.Now().Unix(){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"message":"Token has expired",
		})
	}
 
  e := models.DB.Where("id=?", claims.ID).First(&user).Error
	if e !=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":true,
			"hint":"Error associated with finding user",
			"message":e,
		})
	}

	updateErr := models.DB.Model(&user).Updates(models.User{Verified: true}).Error
	if updateErr !=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":true,
				"hint":"Error associated with updating user",
			"message":updateErr,
		})
	}
  frontend_url := os.Getenv("FRONTEND_URL")
	return c.Redirect(frontend_url)
}

func Login(c *fiber.Ctx) error  {
	
	type LoginInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	err := c.BodyParser(input)

	if err !=nil{
		return c.JSON(fiber.Map{
			"error":true,
			"message":"Please make sure you are sending valid inputs",
		})
	}

	//check if the user exists
	user := new(models.User)

	res := models.DB.Where(&models.User{Email: input.Email}).First(&user)

	if res.RowsAffected <=0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":true,
			"message":"Invalid credentials",
		})
	}

	//compare passwords
	match := helpers.CheckPasswordHash(input.Password,user.Password)
	if !match{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":true,
			"meassage":"Invalid credentials",
		})
	}

	//Check if the user's email has been verified
	if !user.Verified{
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"error":true,
			"meassage":"Email not verified. Please Check the link the mail sent to you and verify your email",
		})
	}

	token, err := helpers.GenerateAccessToken(user.ID)
	if err !=nil{
			return c.JSON(fiber.Map{
			"error":true,
			"message":err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":true,
		"access_token": token,
	})
}