package routes

import (
	"github.com/chukwuka-emi/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  api := app.Group("/api/v1")
	user := api.Group("/user")
	user.Post("/register",controllers.Register )
	user.Get("/verify/:token", controllers.VerifyEmail)
}