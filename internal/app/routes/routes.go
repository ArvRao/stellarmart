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

	userRoutes(app)

	// middleware
	cartRoutes(app)
	return app
}

func userRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", users.UserList)
	userGroup.Get("/:id", users.GetUser)
	userGroup.Get("/signup", users.GetUser)
	userGroup.Get("/login", users.GetUser)
	userGroup.Get("/admin/addproduct", users.GetUser)
	userGroup.Get("/productview", users.GetUser)
	userGroup.Get("/search", users.GetUser)
}

func cartRoutes(app *fiber.App) {
	cartGroup := app.Group("/cart")
	cartGroup.Get("/addtocart", users.GetUser)
	/* cartGroup.Get("/removeitem", cart.removeItem)
	cartGroup.Get("/cartcheckout", cart.buyFromCart)
	cartGroup.Get("/instantBuy", cart.InstantBuy)
	*/
}
