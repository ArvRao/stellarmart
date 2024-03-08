package orders

import (
	"time"

	"github.com/ArvRao/ecommerce-app/internal/app/payment"
	"github.com/ArvRao/ecommerce-app/internal/app/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Order_ID       primitive.ObjectID     `bson:"_id"`
	Order_Cart     []products.ProductUser `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time              `json:"ordered_at" bson:"ordered_at"`
	Price          int                    `json:"total_price" bson:"total_price"`
	Discount       *int                   `json:"discount" bson:"discount"`
	Payment_Method payment.Payment        `json:"payment_method" bson:"payment_method"`
}
