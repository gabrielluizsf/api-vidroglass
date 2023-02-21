package usecase

import (
	"vidroglass/gateway"
	"vidroglass/model"
	"vidroglass/repository"
)

type getCustomerUseCase struct {
}

func NewCustomerUseCase() gateway.CustomerUseCase {
	return &getCustomerUseCase{}
}

func (gc *getCustomerUseCase) GetCustomers() ([]model.Customer, error) {

	clients, err := repository.GetCustomers()

	if err != nil {
		//log
		//todo return error

		return nil, nil

	}
	return clients, nil
}

func (gc *getCustomerUseCase) CreateCustomer(customer model.Customer) (model.Customer, error) {

	customer, err := repository.CreateCustomer(customer)

	if err != nil {
		//log
		//todo return error

		return model.Customer{}, nil

	}
	return customer, nil

}

func (gc *getCustomerUseCase) GetCustomerByID(customer_id int) (model.Customer, error) {

	customer, err := repository.GetCustomerByID(customer_id)

	if err != nil {
		//log
		return model.Customer{}, nil

	}
	return customer, nil

}
