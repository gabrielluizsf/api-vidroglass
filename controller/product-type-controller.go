package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerProductType struct {
	service interfaces.ProductTypeService
}

func NewProductTypeController(service interfaces.ProductTypeService) interfaces.ProductTypeController {
	return &controllerProductType{
		service: service,
	}
}

func (c *controllerProductType) GetProductType(ctx *gin.Context) {
	clientes, err := c.service.GetProductType()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, clientes)

}

func (c *controllerProductType) CreateProductType(ctx *gin.Context) {
	var product_type models.ProductType

	if err := ctx.BindJSON(&product_type); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}
	id, err := c.service.CreateProductType(product_type)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	product_type.Id_tipo = id
	response := models.GoodResponseProductType{"Objeto criado", "Ok", product_type}

	ctx.JSON(200, response)

}

func (c *controllerProductType) GetProductTypeByID(ctx *gin.Context) {
	id_tipo := ctx.Param("id_tipo")
	teste, err := strconv.Atoi(id_tipo)
	product_type, err := c.service.GetProductTypeByID(teste)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Tipo Produto não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, product_type)

}

func (c *controllerProductType) UpdateProductType(ctx *gin.Context) {
	var product_type models.ProductType

	if err := ctx.BindJSON(&product_type); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdateProductType(product_type)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseProductType{"Objeto atualizado", "Ok", product_type}

	ctx.JSON(200, response)

}

func (c *controllerProductType) DeleteProductTypeByID(ctx *gin.Context) {
	id_tipo := ctx.Param("id_tipo")
	teste, err := strconv.Atoi(id_tipo)
	err = c.service.DeleteProductTypeByID(teste)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	response := models.GoodResponseProductType{
		Message:     "Objeto excluido",
		Status:      "Ok",
		ProductType: models.ProductType{}}

	ctx.JSON(200, response)

}
