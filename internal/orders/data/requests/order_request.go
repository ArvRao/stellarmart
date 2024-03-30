package requests

type CreateOrderRequest struct {
	OrderSum uint `validate:"required,min=2,max=1000000" json:"ordersum"`
	UserID   uint `validate:"required" json:"userid"`
}
