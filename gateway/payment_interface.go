package gateway

import "vidroglass/model"

type PaymentUseCase interface {
	GetPayments() ([]model.PaymentMethod, error)
	CreatePayment(model.PaymentMethod) (model.PaymentMethod, error)
	GetPaymentByID(int) (model.PaymentMethod, error)
}
