package Users

import (
	"fmt"
	"strconv"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/users/data/requests"
	"github.com/ArvRao/ecommerce-app/internal/users/data/responses"
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	service "github.com/ArvRao/ecommerce-app/internal/users/services"
	"github.com/ArvRao/ecommerce-app/pkg/database"
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
	createUserRequest := requests.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	helper.ErrorPanic(err)

	userName := controllers.userService.Create(createUserRequest)

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: fmt.Sprintf("Successfully created new user %v", userName),
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controllers *UserController) Update(ctx *fiber.Ctx) error {
	updateUserRequest := requests.UpdateUserRequest{}
	err := ctx.BodyParser(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	err = controllers.userService.Update(id, updateUserRequest.Name)
	helper.ErrorPanic(err)

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Updated users",
		Data:    updateUserRequest.Name,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controllers *UserController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controllers.userService.Delete(id)
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Deleted user",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
}

func (controllers *UserController) GetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	user := controllers.userService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "User fetched successfully",
		Data:    user,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controllers *UserController) FindByAll(ctx *fiber.Ctx) error {
	userResponse := controllers.userService.FindAll()
	webResponse := responses.Response{
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
	users := []Users.User{}
	database.DB.Db.Find(&users)
	return c.Status(200).JSON(users)
}
