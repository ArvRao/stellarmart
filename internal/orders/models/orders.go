package Orders

import (
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderSum uint       `json:"ordersum"`
	UserID   uint       `gorm:"not null"`
	User     Users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
