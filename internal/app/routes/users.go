package routes

import (
	"github.com/ArvRao/ecommerce-app/internal/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// userController := &UserController{}
	userGroup := app.Group("/users")
	userGroup.Get("/signup", controllers.GetUser)
	userGroup.Post("/register", controllers.CreateUser)
	userGroup.Get("/", controllers.UsersList)
	// userGroup.Get("/:id", GetUser)
	// userGroup.Get("/login", GetUser)
	// userGroup.Get("/admin/addproduct", GetUser)
	// userGroup.Get("/productview", GetUser)
	// userGroup.Get("/search", GetUser)
}
