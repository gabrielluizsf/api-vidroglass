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
	DeletePaymentByID(ctx *gin.Context)
}

type PaymentFormService interface {
	GetPaymentFormByID(id_cliente int) (models.PaymentForm, error)
	GetPaymentForm() ([]models.PaymentForm, error)
	CreatePaymentForm(models.PaymentForm) (int, error)
	UpdatePaymentForm(models.PaymentForm) error
	DeletePaymentByID(id_payment int) error
}

type AddressController interface {
	GetAddressByID(ctx *gin.Context)
	GetAddress(ctx *gin.Context)
	CreateAddress(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
	DeleteAddressByID(ctx *gin.Context)
}
type AddressService interface {
	GetAddressByID(id_address int) (models.Address, error)
	GetAddress() ([]models.Address, error)
	CreateAddress(models.Address) (int, error)
	UpdateAddress(models.Address) error
	DeleteAddress(id_address int) error
}

type ProductTypeController interface {
	GetProductTypeByID(ctx *gin.Context)
	GetProductType(ctx *gin.Context)
	CreateProductType(ctx *gin.Context)
	UpdateProductType(ctx *gin.Context)
}

type ProductTypeService interface {
	GetProductTypeByID(id_product_type int) (models.ProductType, error)
	GetProductType() ([]models.ProductType, error)
	CreateProductType(models.ProductType) (int, error)
	UpdateProductType(models.ProductType) error
}

type ProductController interface {
	GetProductByID(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

type ProductService interface {
	GetProductByID(id_product int) (models.Product, error)
	GetProduct() ([]models.ProductPayload, error)
	CreateProduct(models.Product) (int, error)
	UpdateProduct(models.Product) error
}

type ItemController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	UpdateItemById(ctx *gin.Context)
	GetItemById(ctx *gin.Context)
}

type ItemService interface {
	GetItemById(id_cliente int) (models.Item, error)
	FindAll() ([]models.Item, error)
	Save(models.Item) (int, error)
	UpdateItemById(models.Item) error
	GetInvoiceItens(id_nota int) ([]models.ItemPayload, error)
}

type NotaController interface {
	GetNotaByID(ctx *gin.Context)
	GetNota(ctx *gin.Context)
	CreateNota(ctx *gin.Context)
	UpdateNota(ctx *gin.Context)
}

type NotaService interface {
	GetNotaByID(id_nota int) (models.NotaPayload, error)
	GetNota() ([]models.Nota, error)
	CreateNota() (int, error)
	UpdateNota(models.Nota) (models.Nota, error)
}
