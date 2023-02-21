package gateway

import "vidroglass/model"

type ProductUseCase interface {
	GetProducts() ([]model.ProductObject, error)
	CreateProduct(model.Product) (model.Product, error)
	GetProductByID(int) (model.ProductObject, error)
	GetProductTypes() ([]model.ProductType, error)
	CreateProductType(model.ProductType) (model.ProductType, error)
}
