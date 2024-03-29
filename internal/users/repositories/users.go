package Users

import (
	"errors"

	"github.com/ArvRao/ecommerce-app/internal/helper"
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	"github.com/ArvRao/ecommerce-app/pkg/database"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

type UserRepository interface {
	Save(Users.User)
	Update(id int, name string) error
	Delete(userId int)
	FindById(userId int) (Users.User, error)
	FindAll() []Users.User
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user Users.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user Users.User) {
	database.DB.Db.Create(&user)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(id int, name string) error {
	result := u.Db.Model(&Users.User{}).Where("id = ?", id).Update("name", name)
	return result.Error
}

func (u *UserRepositoryImpl) FindById(userId int) (Users.User, error) {
	var user Users.User
	// result := u.Db.Find(&user, userId)
	result := u.Db.First(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (u *UserRepositoryImpl) FindAll() []Users.User {
	users := []Users.User{}
	database.DB.Db.Find(&users)
	return users
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
