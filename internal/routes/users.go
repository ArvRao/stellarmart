package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(userController *controllers.UserController) *fiber.App {
	router := fiber.New()

	router.Route("/users", func(router fiber.Router) {
		router.Post("/", userController.Create)
		router.Get("/", userController.FindByAll)
	})
	return router

}

func UserRoutes(app *fiber.App) {
	// userController * controllers.UserController{}
	userGroup := app.Group("/users")
	userGroup.Get("/signup", controllers.GetUser)
	// userGroup.Post("/register", controllers.Create)
	userGroup.Get("/", controllers.UsersList)
	// userGroup.Get("/:id", GetUser)
	// userGroup.Get("/login", GetUser)
	// userGroup.Get("/admin/addproduct", GetUser)
	// userGroup.Get("/productview", GetUser)
	// userGroup.Get("/search", GetUser)
}
