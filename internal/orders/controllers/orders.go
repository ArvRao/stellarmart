package Orders

import (
	"fmt"
	"strconv"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/orders/data/requests"
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/services"
	"github.com/ArvRao/ecommerce-app/internal/users/data/responses"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService Orders.OrderService
}

func NewOrderController(service Orders.OrderService) *OrderController {
	return &OrderController{orderService: service}
}

func (controllers *OrderController) Create(ctx *fiber.Ctx) error {
	createOrderRequest := requests.CreateOrderRequest{}
	err := ctx.BodyParser(&createOrderRequest)
	helper.ErrorPanic(err)

	order := controllers.orderService.Create(createOrderRequest)

	// check the number
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: fmt.Sprintf("Successfully created new order %v", order),
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controllers *OrderController) Delete(ctx *fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	helper.ErrorPanic(err)
	controllers.orderService.Delete(id)
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Deleted order",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
}

func (controllers *OrderController) GetById(ctx *fiber.Ctx) error {
	orderId := ctx.Params("orderId")
	id, err := strconv.Atoi(orderId)
	helper.ErrorPanic(err)

	order := controllers.orderService.FindById(id)
	if order.Id == 0 {
		fmt.Println("Given order does not exist")
	}
	helper.ErrorPanic(err)

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Order fetched successfully",
		Data:    order,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controllers *OrderController) FindByAll(ctx *fiber.Ctx) error {
	userResponse := controllers.orderService.FindAll()
	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Got all the orders",
		Data:    userResponse,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(webResponse)
}

// Update controller is remaining
