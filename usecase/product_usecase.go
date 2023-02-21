package usecase

import (
	"vidroglass/gateway"
	"vidroglass/model"
	"vidroglass/repository"
)

type getProductUseCase struct {
}

func NewProductUseCase() gateway.ProductUseCase {
	return &getProductUseCase{}
}

func (gc *getProductUseCase) GetProducts() ([]model.ProductObject, error) {

	products, err := repository.GetProducts()

	if err != nil {
		//log
		return nil, nil

	}
	return products, nil
}

func (gc *getProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	id, err := repository.CreateProduct(product)

	if err != nil {
		//log
		//todo return error

		return model.Product{}, nil

	}
	product.ProductID = id
	return product, nil

}

func (gc *getProductUseCase) GetProductByID(product_id int) (model.ProductObject, error) {

	product, err := repository.GetProductByID(product_id)

	if err != nil {
		//log
		//todo return error
		return model.ProductObject{}, nil

	}
	return product, nil

}
func (gc *getProductUseCase) GetProductTypes() ([]model.ProductType, error) {

	types, err := repository.GetProductTypes()

	if err != nil {
		//log
		return nil, nil

	}
	return types, nil
}

func (gc *getProductUseCase) CreateProductType(productT model.ProductType) (model.ProductType, error) {

	id, err := repository.CreateProductType(productT)

	if err != nil {
		//log
		//todo return error

		return model.ProductType{}, nil

	}
	productT.TypeID = id
	return productT, nil

}
