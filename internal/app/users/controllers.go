package users

import (
	"github.com/ArvRao/ecommerce-app/internal/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	return c.SendString("User ID: " + userId)
}

// UserController is the controller for handling user-related requests
type UserController struct {
	DB *gorm.DB // Assuming you have a GORM database connection
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&user)
	return c.Status(201).JSON(user)
}

func UsersList(c *fiber.Ctx) error {
	users := []User{}
	database.DB.Db.Find(&users)
	return c.Status(200).JSON(users)
}
