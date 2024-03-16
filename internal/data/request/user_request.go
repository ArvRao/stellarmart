package request

type CreateUserRequest struct {
	UserName string `validate:"required,min=2,max=28" json:"user"`
}

type UpdateUserRequest struct {
	Id  int `validate:"required"`
	Age int `validate:"required,max=28,min=2" json:"age"`
}
