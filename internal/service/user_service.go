package service

import (
	"fmt"

	"github.com/ArvRao/ecommerce-app/internal/data/request"
	"github.com/ArvRao/ecommerce-app/internal/data/response"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/models"
	"github.com/ArvRao/ecommerce-app/internal/repository"
	"github.com/go-playground/validator"
)

type UserService interface {
	Create(user request.CreateUserRequest) models.User
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []models.User
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) models.User {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := models.User{
		Name: user.Name,
	}
	u.UserRepository.Save(userModel)
	return userModel
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []models.User {
	fmt.Println("inside user service")
	// work on fixing the repository problem

	result := u.UserRepository.FindAll()
	fmt.Println("after result", result)
	// var users []response.UserResponse

	// users := []models.User{}
	// database.DB.Db.Find(&users)
	// fmt.Println(users)
	return result

}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := response.UserResponse{
		Name: userData.Name,
	}
	return userResponse
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	u.UserRepository.Update(userData)
}
