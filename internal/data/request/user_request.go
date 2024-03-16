package request

type CreateUserRequest struct {
	Name string `validate:"required,min=2,max=28" json:"name"`
}

type UpdateUserRequest struct {
	Id  int `validate:"required"`
	Age int `validate:"required,max=28,min=2" json:"age"`
}
