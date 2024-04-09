package Orders

import (
	"fmt"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/orders/data/requests"
	"github.com/ArvRao/ecommerce-app/internal/orders/data/responses"
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/models"
	repository "github.com/ArvRao/ecommerce-app/internal/orders/repositories"
	"github.com/go-playground/validator"
)

type OrderService interface {
	Create(order requests.CreateOrderRequest) uint
	Update(id int, ordersum int) error
	Delete(orderId int)
	FindById(orderId int) responses.OrderResponse
	FindAll() []Orders.Order
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	validate        *validator.Validate
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		validate:        validate,
	}
}

func (u *OrderServiceImpl) Create(order requests.CreateOrderRequest) uint {
	err := u.validate.Struct(order)
	helper.ErrorPanic(err)
	orderModel := Orders.Order{
		OrderSum: order.OrderSum,
		UserID:   order.UserID,
	}
	u.OrderRepository.Save(orderModel)
	fmt.Println(orderModel.ID)
	// not working here, returning 0
	return orderModel.ID
}

func (u *OrderServiceImpl) Delete(userId int) {
	u.OrderRepository.Delete(userId)
}

func (u *OrderServiceImpl) FindAll() []Orders.Order {
	result := u.OrderRepository.FindAll()
	return result
}

func (u *OrderServiceImpl) FindById(userId int) responses.OrderResponse {
	orderData, err := u.OrderRepository.FindById(userId)
	// orderData is empty check
	helper.ErrorPanic(err)
	fmt.Println(orderData.OrderSum)
	orderResponse := responses.OrderResponse{
		Id:       int(orderData.ID),
		OrderSum: int(orderData.OrderSum),
	}
	return orderResponse
}

func (u *OrderServiceImpl) Update(id int, ordersum int) error {
	_, err := u.OrderRepository.FindById(id)
	if err != nil {
		return err
	}
	return u.OrderRepository.Update(id, ordersum)
}