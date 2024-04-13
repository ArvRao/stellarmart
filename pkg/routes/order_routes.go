package routes

import (
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(api fiber.Router, orderController *Orders.OrderController) {
	userRoutes := api.Group("/orders")
	userRoutes.Get("/all", orderController.FindByAll)
	userRoutes.Post("/", orderController.Create)
	userRoutes.Delete("/:orderId", orderController.Delete)
	// userRoutes.Patch("/:userId", orderController.Update)
	// GET /api/orders?orderId=3
	userRoutes.Get("/", orderController.GetById)
}
