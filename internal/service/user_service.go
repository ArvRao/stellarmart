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
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() response.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := models.User{
		Name: user.UserName,
	}
	u.UserRepository.Save(userModel)
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()
	var users []response.UserResponse

	for _, value := range result {
		user := response.UserResponse{
			Id:   int(value.ID),
			Name: value.Name,
		}
		users = append(users, user)
	}
	return users
}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := response.UserResponse{
		Id:   int(userData.ID),
		Name: userData.Name,
	}
	return userResponse
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Age = uint16(user.Age)
	u.UserRepository.Update(userData)
}
