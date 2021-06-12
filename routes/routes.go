package routes

import (
	"github.com/chukwuka-emi/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	user := api.Group("/user")
	user.Post("/register", controllers.Register)
	user.Get("/verify/:token", controllers.VerifyEmail)
	user.Post("/login", controllers.Login)
}

func SecuredRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	account := api.Group("/account")
	account.Post("/open", controllers.OpenAccount)
	account.Get("/:id", controllers.FindAccount)
	account.Delete("/:account_number", controllers.CloseAccount)
	account.Post("/reactivate/:account_number", controllers.ReactivateAccount)
}
