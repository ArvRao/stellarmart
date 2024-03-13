package users

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// userController := &UserController{}
	userGroup := app.Group("/users")
	userGroup.Get("/signup", GetUser)
	userGroup.Post("/register", CreateUser)
	userGroup.Get("/", UsersList)
	// userGroup.Get("/:id", GetUser)
	// userGroup.Get("/login", GetUser)
	// userGroup.Get("/admin/addproduct", GetUser)
	// userGroup.Get("/productview", GetUser)
	// userGroup.Get("/search", GetUser)
}
