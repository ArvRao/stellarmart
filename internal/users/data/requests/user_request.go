package requests

type CreateUserRequest struct {
	Name      string `validate:"required,min=2,max=28" json:"name"`
	FirstName string `json:"first_name" validate:"required,min=3,max=25" gorm:"text;not null;"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50" gorm:"text;not null;"`
	Email     string `gorm:"email" json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Role      string `json:"role" validate:"required,role"`
	Address   string `json:"address" validate:"required,min=12,max=60" gorm:"text;not null;"`
	City      string `json:"city" validate:"required" gorm:"text;not null;"`
	State     string `json:"state" validate:"required" gorm:"text;not null;"`
	ZipCode   string `json:"zipcode" validate:"required,min=6,max=6" gorm:"text;not null;"`
}

type UpdateUserRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=2,max=28" json:"name"`
}
