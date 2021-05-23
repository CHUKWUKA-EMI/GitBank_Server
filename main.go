package main

import (
	"log"
	"os"

	// "github.com/chukwuka-emi/helpers"
	"github.com/chukwuka-emi/models"
	"github.com/chukwuka-emi/routes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

type Claims struct {
 jwt.StandardClaims
 ID uint `gorm:"primaryKey"`
 }

func main(){
	app := fiber.New()

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	app.Use(cors.New())

	//free routes
	routes.AuthRoutes(app)

	//Secure Routes with JWT Auth middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Claims: Claims{},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":e.Error(),
			})
		},
	}))
	
	
	db_url:= os.Getenv("DB_URL")
	models.ConnectDB(db_url)
	
  app.Listen(":3000")
}