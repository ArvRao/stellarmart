// internal/app/routes/routes.go

package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the routes for the application using Go Fiber.
// setup major routes for the application. Routes from here to other services

func cartRoutes(app *fiber.App) {
	cartGroup := app.Group("/cart")
	cartGroup.Get("/addtocart", controllers.GetUser)
	/* cartGroup.Get("/removeitem", cart.removeItem)
	cartGroup.Get("/cartcheckout", cart.buyFromCart)
	cartGroup.Get("/instantBuy", cart.InstantBuy)
	*/
}
