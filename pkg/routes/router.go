package routes

import (
	Users "github.com/ArvRao/ecommerce-app/internal/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(userController *Users.UserController) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	// user routes
	SetupUserRoutes(api, userController)

	return app
}
