package routes

import (
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/controllers"
	Users "github.com/ArvRao/ecommerce-app/internal/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(userController *Users.UserController, orderController *Orders.OrderController) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	// user routes
	SetupUserRoutes(api, userController)
	// order routes
	SetupOrderRoutes(api, orderController)
	// cart routes

	return app
}
