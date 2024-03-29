package Users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=2,max=100" gorm:"text;not null;"`
	// Orders []Orders.Order `gorm:"foreignKey:UserID"`
}

type Address struct {
	gorm.Model
	HouseName string `json:"housename"`
	City      string `json:"city"`
	UserID    uint
}
