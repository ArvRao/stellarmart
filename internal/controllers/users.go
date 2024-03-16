package controllers

import (
	"fmt"
	"strconv"

	"github.com/ArvRao/ecommerce-app/internal/data/request"
	"github.com/ArvRao/ecommerce-app/internal/data/response"
	"github.com/ArvRao/ecommerce-app/internal/database"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/models"
	"github.com/ArvRao/ecommerce-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	return c.SendString("User ID: " + userId)
}

// UserController is the controller for handling user-related requests
type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (controllers *UserController) Create(ctx *fiber.Ctx) error {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	helper.ErrorPanic(err)

	userName := controllers.userService.Create(createUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: fmt.Sprintf("Successfully created new user %v", userName),
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controllers *UserController) Update(ctx *fiber.Ctx) error {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controllers.userService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Updated users",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controllers *UserController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controllers.userService.Delete(id)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Deleted user",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
}

/*
	func (controllers *UserController) FindById(ctx *fiber.Ctx) error {

}
*/
func (controllers *UserController) FindByAll(ctx *fiber.Ctx) error {
	userResponse := controllers.userService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Got all the users",
		Data:    userResponse,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
}

// func CreateUser(c *fiber.Ctx) error {
// 	user := new(models.User)
// 	if err := c.BodyParser(user); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	database.DB.Db.Create(&user)
// 	return c.Status(201).JSON(user)
// }

func UsersList(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find(&users)
	return c.Status(200).JSON(users)
}
