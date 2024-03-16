// internal/app/routes/routes.go

package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the application using Go Fiber.
// setup major routes for the application. Routes from here to other services
func SetupRoutes() *fiber.App {
	app := fiber.New()

	// Define routes and associate them with handlers
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Status":  "success",
			"message": "welcome to stellarmart",
		})
	})

	// route to routes -> users
	UserRoutes(app)

	// middleware
	cartRoutes(app)
	return app
}

func cartRoutes(app *fiber.App) {
	cartGroup := app.Group("/cart")
	cartGroup.Get("/addtocart", controllers.GetUser)
	/* cartGroup.Get("/removeitem", cart.removeItem)
	cartGroup.Get("/cartcheckout", cart.buyFromCart)
	cartGroup.Get("/instantBuy", cart.InstantBuy)
	*/
}
