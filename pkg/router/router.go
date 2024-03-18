package router

import (
	"github.com/ArvRao/ecommerce-app/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(userController *controllers.UserController) *fiber.App {
	router := fiber.New()

	router.Route("/users", func(router fiber.Router) {
		router.Post("/", userController.Create)
		router.Get("/", userController.FindByAll)
		router.Delete("/:userId", userController.Delete)
		router.Patch("/:userId", userController.Update)
		router.Get("/:userId", userController.GetById)
	})
	return router

}
