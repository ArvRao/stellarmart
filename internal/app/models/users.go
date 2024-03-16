package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=2,max=100" gorm:"text;not null;"`
	Age  uint16 `json:"age"`
}
