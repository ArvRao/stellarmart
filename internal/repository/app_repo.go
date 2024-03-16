package repository

import (
	"fmt"

	"github.com/ArvRao/ecommerce-app/internal/app/models"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user models.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
	fmt.Println(result)
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(models.User) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(models.User) {
	panic("unimplemented")
}

func (u *UserRepositoryImpl) FindAll() []models.User {
	var user []models.User
	result := u.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	fmt.Println(result)
	return user
}

type UserRepository interface {
	Save(models.User)
	Update(models.User)
	Delete(userId int)
	FindAll() []models.User
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
