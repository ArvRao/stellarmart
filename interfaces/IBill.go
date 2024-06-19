package interfaces

import "github.com/ArvRao/ecommerce-app/model"

type IBillGenerator interface {
	GenerateInvoice(model.Invoice) string
}
