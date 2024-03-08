package users

import (
	"time"

	"github.com/ArvRao/ecommerce-app/internal/app/orders"
	"github.com/ArvRao/ecommerce-app/internal/app/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID     `json:"_id" bson:"_id"`
	First_Name      *string                `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string                `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string                `json:"password"  validate:"required,min=8"`
	Email           *string                `json:"email"  validate:"email, required"`
	Phone           *string                `json:"phone" validate:"required"`
	Token           *string                `json:"token"`
	Refresh_Token   *string                `json:"refresh_token"`
	Created_At      time.Time              `json:"created_at"`
	Updated_At      time.Time              `json:"updated_at"`
	User_ID         string                 `json:"user_id"`
	UserCart        []products.ProductUser `json:"usercart" bson:"usercart"`
	Address_Details []Address              `json:"address" bson:"address"`
	Order_Status    []orders.Order         `json:"orders" bson:"orders"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pincode" bson:"pincode"`
}
