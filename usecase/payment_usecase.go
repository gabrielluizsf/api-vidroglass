package usecase

import (
	"vidroglass/gateway"
	"vidroglass/model"
	"vidroglass/repository"
)

type getPaymentUseCase struct {
}

func NewPaymentUseCase() gateway.PaymentUseCase {
	return &getPaymentUseCase{}
}

func (gc *getPaymentUseCase) GetPayments() ([]model.PaymentMethod, error) {

	clients, err := repository.GetPaymentsMethod()

	if err != nil {
		//log
		//todo return error

		return nil, nil

	}
	return clients, nil
}

func (gc *getPaymentUseCase) CreatePayment(payment model.PaymentMethod) (model.PaymentMethod, error) {

	id, err := repository.CreatePaymentMethod(payment)

	if err != nil {
		//log
		//todo return error

		return model.PaymentMethod{}, nil

	}
	payment.PaymentID = id
	return payment, nil

}

func (gc *getPaymentUseCase) GetPaymentByID(payment_id int) (model.PaymentMethod, error) {

	payment, err := repository.GetPaymentMethodByID(payment_id)

	if err != nil {
		//log
		return model.PaymentMethod{}, nil

	}
	return payment, nil

}
