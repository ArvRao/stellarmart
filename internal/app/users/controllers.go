package users

import "github.com/gofiber/fiber/v2"

func UserList(c *fiber.Ctx) error {
	return c.SendString("User List")
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	return c.SendString("User ID: " + userId)
}
