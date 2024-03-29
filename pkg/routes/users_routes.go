package routes

import (
	Users "github.com/ArvRao/ecommerce-app/internal/users/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, userController *Users.UserController) {
	userRoutes := api.Group("/users")
	userRoutes.Get("/", userController.FindByAll)
	userRoutes.Post("/", userController.Create)
	userRoutes.Delete("/:userId", userController.Delete)
	userRoutes.Patch("/:userId", userController.Update)
	userRoutes.Get("/:userId", userController.GetById)
}
