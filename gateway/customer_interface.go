package gateway

import "vidroglass/model"

type CustomerUseCase interface {
	GetCustomers() ([]model.Customer, error)
	CreateCustomer(model.Customer) (model.Customer, error)
	GetCustomerByID(int) (model.Customer, error)
}
