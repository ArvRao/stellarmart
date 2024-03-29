package Users

import (
	"github.com/ArvRao/ecommerce-app/internal/helper"
	"github.com/ArvRao/ecommerce-app/internal/users/data/requests"
	"github.com/ArvRao/ecommerce-app/internal/users/data/responses"
	Users "github.com/ArvRao/ecommerce-app/internal/users/models"
	repository "github.com/ArvRao/ecommerce-app/internal/users/repositories"

	"github.com/go-playground/validator"
)

type UserService interface {
	Create(user requests.CreateUserRequest) string
	Update(id int, name string) error
	Delete(userId int)
	FindById(userId int) responses.UserResponse
	FindAll() []Users.User
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

func (u *UserServiceImpl) Create(user requests.CreateUserRequest) string {
	err := u.validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := Users.User{
		Name: user.Name,
	}
	u.UserRepository.Save(userModel)
	return userModel.Name
}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []Users.User {
	result := u.UserRepository.FindAll()
	return result
}

func (u *UserServiceImpl) FindById(userId int) responses.UserResponse {
	userData, err := u.UserRepository.FindById(userId)
	helper.ErrorPanic(err)
	userResponse := responses.UserResponse{
		Name: userData.Name,
	}
	return userResponse
}

func (u *UserServiceImpl) Update(id int, name string) error {
	_, err := u.UserRepository.FindById(id)
	if err != nil {
		return err
	}
	return u.UserRepository.Update(id, name)
}
