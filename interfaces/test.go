package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/models"
)

type ClienteController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	UpdateClientById(ctx *gin.Context)
	GetClientById(ctx *gin.Context)
}

type ClienteService interface {
	GetClientById(id_cliente int) (models.Cliente, error)
	FindAll() ([]models.Cliente, error)
	Save(models.Cliente) (int, error)
	UpdateClientById(models.Cliente) error
}

type PaymentFormController interface {
	GetPaymentFormByID(ctx *gin.Context)
	GetPaymentForm(ctx *gin.Context)
	CreatePaymentForm(ctx *gin.Context)
	UpdatePaymentForm(ctx *gin.Context)
}

type PaymentFormService interface {
	GetPaymentFormByID(id_cliente int) (models.PaymentForm, error)
	GetPaymentForm() ([]models.PaymentForm, error)
	CreatePaymentForm(models.PaymentForm) (int, error)
	UpdatePaymentForm(models.PaymentForm) error
}

type AddressController interface {
	GetAddressByID(ctx *gin.Context)
	GetAddress(ctx *gin.Context)
	CreateAddress(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
}
type AddressService interface {
	GetAddressByID(id_address int) (models.Address, error)
	GetAddress() ([]models.Address, error)
	CreateAddress(models.Address) (int, error)
	UpdateAddress(models.Address) error
}
