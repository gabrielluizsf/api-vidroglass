package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerProduct struct {
	service interfaces.ProductService
}

func NewProductController(service interfaces.ProductService) interfaces.ProductController {
	return &controllerProduct{
		service: service,
	}
}

func (c *controllerProduct) GetProduct(ctx *gin.Context) {
	produto, err := c.service.GetProduct()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, produto)

}

func (c *controllerProduct) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}
	id, err := c.service.CreateProduct(product)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	product.Id_produto = id
	response := models.GoodResponseProduct{"Objeto criado", "Ok", product}

	ctx.JSON(200, response)

}

func (c *controllerProduct) GetProductByID(ctx *gin.Context) {
	id_produto := ctx.Param("id_tipo")
	id_produtostr, err := strconv.Atoi(id_produto)
	produto, err := c.service.GetProductByID(id_produtostr)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Produto não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, produto)

}

func (c *controllerProduct) UpdateProduct(ctx *gin.Context) {
	var produto models.Product

	if err := ctx.BindJSON(&produto); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdateProduct(produto)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseProduct{"Objeto atualizado", "Ok", produto}

	ctx.JSON(200, response)

}

func (c *controllerProduct) DeleteProductById(ctx *gin.Context) {
	id_produto := ctx.Param("id_product")
	id_produtostr, err := strconv.Atoi(id_produto)
	err = c.service.DeleteProductById(id_produtostr)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Produto não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	response := models.GoodResponseProduct{"Objeto excluido", "Ok", models.Product{}}

	ctx.JSON(200, response)

}
