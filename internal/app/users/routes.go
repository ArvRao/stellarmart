package users

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", UserList)
	userGroup.Get("/:id", GetUser)
	userGroup.Get("/signup", GetUser)
	userGroup.Get("/login", GetUser)
	userGroup.Get("/admin/addproduct", GetUser)
	userGroup.Get("/productview", GetUser)
	userGroup.Get("/search", GetUser)
}
