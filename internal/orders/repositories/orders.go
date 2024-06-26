package Orders

import (
	"errors"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/orders/data/responses"
	Orders "github.com/ArvRao/ecommerce-app/internal/orders/models"
	"github.com/ArvRao/ecommerce-app/pkg/database"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements OrderRepository.
func (o *OrderRepositoryImpl) Delete(orderId int) {
	var order Orders.Order
	result := o.Db.Where("id = ?", orderId).Delete(&order)
	helper.ErrorPanic(result.Error)
}

// FindAll implements OrderRepository.
func (o *OrderRepositoryImpl) FindAll() []responses.OrderResponse {
	orders := []Orders.Order{}
	database.DB.Db.Find(&orders)
	var orderResponses []responses.OrderResponse
	for _, order := range orders {
		orderResponse := responses.OrderResponse{
			OrderSum:  int(order.OrderSum),
			Id:        int(order.UserID),
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			UserID:    order.UserID,
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	return orderResponses
}

// FindById implements OrderRepository.
func (o *OrderRepositoryImpl) FindById(orderId int) (Orders.Order, error) {
	var order Orders.Order
	// result := o.Db.Find(&order, orderId)
	result := o.Db.First(&order, orderId)
	if result != nil {
		return order, nil
	} else {
		return order, errors.New("order is not found")
	}
}

// Save implements OrderRepository.
func (o *OrderRepositoryImpl) Save(order Orders.Order) {

	database.DB.Db.Create(&order)
}

// Update implements OrderRepository.
func (o *OrderRepositoryImpl) Update(id int, ordersum int) error {
	result := o.Db.Model(&Orders.Order{}).Where("id = ?", ordersum).Update("ordersum", ordersum)
	return result.Error
}

type OrderRepository interface {
	Save(Orders.Order)
	Update(id int, ordersum int) error
	Delete(orderId int)
	FindById(orderId int) (Orders.Order, error)
	FindAll() []responses.OrderResponse
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{Db: db}
}
