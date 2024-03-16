package repository

import (
	"fmt"

	"github.com/ArvRao/ecommerce-app/internal/app/users"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user users.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
	fmt.Println(result)
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(users.User) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(users.User) {
	panic("unimplemented")
}

func (u *UserRepositoryImpl) FindAll() []users.User {
	var user []users.User
	result := u.Db.Find(&user)
	helper.ErrorPanic(result.Error)
	fmt.Println(result)
	return user
}

type UserRepository interface {
	Save(users.User)
	Update(users.User)
	Delete(userId int)
	FindAll() []users.User
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
