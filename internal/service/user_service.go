package service

import (
	"github.com/ArvRao/ecommerce-app/internal/data/request"
	"github.com/ArvRao/ecommerce-app/internal/data/response"
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/models"
	"github.com/ArvRao/ecommerce-app/internal/repository"
	"github.com/go-playground/validator"
)

type UserService interface {
	Create(user request.CreateUserRequest) string
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

func (u *UserServiceImpl) Create(user request.CreateUserRequest) string {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := models.User{
		Name: user.Name,
	}
	u.UserRepository.Save(userModel)
	return userModel.Name
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []models.User {
	result := u.UserRepository.FindAll()
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
