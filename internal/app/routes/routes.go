// internal/app/routes/routes.go

package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/app/users"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the application using Go Fiber.
// setup major routes for the application
func SetupRoutes() *fiber.App {
	app := fiber.New()

	// Define routes and associate them with handlers
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Pong!")
	})

	users.UserRoutes(app)

	// middleware
	cartRoutes(app)
	return app
}

func cartRoutes(app *fiber.App) {
	cartGroup := app.Group("/cart")
	cartGroup.Get("/addtocart", users.GetUser)
	/* cartGroup.Get("/removeitem", cart.removeItem)
	cartGroup.Get("/cartcheckout", cart.buyFromCart)
	cartGroup.Get("/instantBuy", cart.InstantBuy)
	*/
}
