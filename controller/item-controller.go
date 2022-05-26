package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerItem struct {
	service interfaces.ItemService
}

func NewItemController(service interfaces.ItemService) interfaces.ItemController {
	return &controllerItem{
		service: service,
	}
}

func (c *controllerItem) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	if len(items) <= 0 {
		ctx.JSON(400, "Nada encontrado")
		return

	}

	ctx.JSON(200, items)

}

func (c *controllerItem) Save(ctx *gin.Context) {
	var item models.Item

	if err := ctx.BindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}
	id, err := c.service.Save(item)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	item.Id_item = id
	response := models.GoodResponseItem{"Objeto criado", "Ok", item}

	ctx.JSON(200, response)

}

func (c *controllerItem) GetItemById(ctx *gin.Context) {
	id_item := ctx.Param("id_item")
	teste, err := strconv.Atoi(id_item)
	customer, err := c.service.GetItemById(teste)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Item não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)

}

func (c *controllerItem) UpdateItemById(ctx *gin.Context) {
	var item models.Item

	if err := ctx.BindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdateItemById(item)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseItem{"Objeto atualizado", "Ok", item}

	ctx.JSON(200, response)

}
