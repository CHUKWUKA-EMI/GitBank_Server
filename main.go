package main

import (
	"log"
	"os"

	"github.com/chukwuka-emi/models"
	"github.com/chukwuka-emi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)


func main(){
	app := fiber.New()

	
	app.Use(cors.New())
	
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	routes.SetupRoutes(app)

	db_url:= os.Getenv("DB_URL")
	models.ConnectDB(db_url)
	
  app.Listen(":3000")
}