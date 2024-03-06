// internal/app/routes/routes.go

package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the application using Go Fiber.
func SetupRoutes() *fiber.App {
	app := fiber.New()

	// Define routes and associate them with handlers
	app.Get("/hello", handlers.HelloHandler)

	return app
}
