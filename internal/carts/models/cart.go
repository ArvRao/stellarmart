package Cart

import (
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"uniqueIndex"`
	Status string     `gorm:"default:'pending'"`
	User   Users.User `gorm:"foreignKey:UserID;references:ID"`
}
