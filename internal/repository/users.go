package repository

import (
	"errors"
	"fmt"

	"github.com/ArvRao/ecommerce-app/internal/data/request"
	"github.com/ArvRao/ecommerce-app/internal/database"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/models"
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
func (u *UserRepositoryImpl) Save(user models.User) {
	database.DB.Db.Create(&user)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user models.User) {
	var updateUser = request.UpdateUserRequest{
		Id: int(user.ID),
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)

}

func (u *UserRepositoryImpl) FindById(userId int) (models.User, error) {
	var user models.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (u *UserRepositoryImpl) FindAll() []models.User {
	users := []models.User{}
	database.DB.Db.Find(&users)
	return users
}

type UserRepository interface {
	Save(models.User)
	Update(models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}
