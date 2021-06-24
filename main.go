package main

import (
	"log"
	"os"

	// "github.com/chukwuka-emi/helpers"
	"github.com/chukwuka-emi/models"
	"github.com/chukwuka-emi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Use(cors.New())

	//free routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "WELCOME TO GITBANK SERVER...",
		})
	})
	routes.AuthRoutes(app)

	//Secure Routes with JWT Auth middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": e.Error(),
			})
		},
	}))

	//secured routes
	routes.SecuredRoutes(app)

	db_url := os.Getenv("DATABASE_URL")
	models.ConnectDB(db_url)

	app.Listen(":3000")
}
