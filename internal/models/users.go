package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=2,max=100" gorm:"text;not null;"`
	Age  uint16 `json:"age"`
}

type Address struct {
	gorm.Model
	HouseName string `json:"housename"`
	City      string `json:"city"`
	UserID    uint
}
