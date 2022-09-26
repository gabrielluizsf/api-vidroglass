package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerCliente struct {
	service interfaces.ClienteService
}

func NewClienteController(service interfaces.ClienteService) interfaces.ClienteController {
	return &controllerCliente{
		service: service,
	}
}

func (c *controllerCliente) FindAll(ctx *gin.Context) {
	clientes, err := c.service.FindAll()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, clientes)

}

func (c *controllerCliente) Save(ctx *gin.Context) {
	var cliente models.Cliente

	if err := ctx.BindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"A requisição está incorreta!", "Erro na requisição", err.Error()})
		return
	}
	id, err := c.service.Save(cliente)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(500, response)
		return
	}
	cliente.Id_cliente = id
	response := models.GoodResponse{"Objeto criado", "Ok", cliente}

	ctx.JSON(200, response)

}

func (c *controllerCliente) GetClientById(ctx *gin.Context) {
	id_cliente := ctx.Param("id_cliente")
	teste, err := strconv.Atoi(id_cliente)
	customer, err := c.service.GetClientById(teste)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Cliente não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)

}

func (c *controllerCliente) UpdateClientById(ctx *gin.Context) {
	var cliente models.Cliente

	if err := ctx.BindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdateClientById(cliente)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponse{"Objeto atualizado", "Ok", cliente}

	ctx.JSON(200, response)

}
