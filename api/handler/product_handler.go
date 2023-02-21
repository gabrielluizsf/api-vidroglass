package handler

import (
	"fmt"
	"net/http"
	"vidroglass/api/util"
	"vidroglass/gateway"

	"vidroglass/model"

	"github.com/gin-gonic/gin"
)

type ProductHandlerInterface interface {
	GetProductsHandler(ctx *gin.Context)
	CreateProductHandler(ctx *gin.Context)
	GetProductByIDHandler(ctx *gin.Context)
	GetProductTypesHandler(ctx *gin.Context)
	CreateProductTypeHandler(ctx *gin.Context)
}

type productHandler struct {
	productUseCase gateway.ProductUseCase
}

func NewProductHandler(usecase gateway.ProductUseCase) ProductHandlerInterface {
	return &productHandler{
		productUseCase: usecase,
	}
}

func (c *productHandler) GetProductsHandler(ctx *gin.Context) {

	response, err := c.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "response")

	}
	ctx.JSON(200, response)
}

func (c *productHandler) CreateProductHandler(ctx *gin.Context) {
	var product model.Product

	if err := ctx.BindJSON(&product); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := c.productUseCase.CreateProduct(product)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  product,
	}

	ctx.JSON(200, response)
}

func (c *productHandler) GetProductByIDHandler(ctx *gin.Context) {

	productId, err := getProductID(ctx)
	if err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "product_id inválido!")
		ctx.JSON(http.StatusBadRequest, response)
	}

	product, err := c.productUseCase.GetProductByID(productId)

	if err != nil {
		response := util.BuildErrorResponse("Product not found!", "not found")
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, product)
}

func (c *productHandler) GetProductTypesHandler(ctx *gin.Context) {

	response, err := c.productUseCase.GetProductTypes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "response")

	}
	ctx.JSON(200, response)
}

func (c *productHandler) CreateProductTypeHandler(ctx *gin.Context) {
	var productType model.ProductType

	if err := ctx.BindJSON(&productType); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	productType, err := c.productUseCase.CreateProductType(productType)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  productType,
	}

	ctx.JSON(200, response)
}
