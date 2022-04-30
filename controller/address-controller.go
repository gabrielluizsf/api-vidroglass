package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerAddress struct {
	service interfaces.AddressService
}

func NewAddressController(service interfaces.AddressService) interfaces.AddressController {
	return &controllerAddress{
		service: service,
	}
}

func (c *controllerAddress) GetAddress(ctx *gin.Context) {
	address, err := c.service.GetAddress()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, address)

}

func (c *controllerAddress) CreateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.BindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}
	id, err := c.service.CreateAddress(address)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	address.Id_address = id
	response := models.GoodResponseAddress{"Objeto criado", "Ok", address}

	ctx.JSON(200, response)

}

func (c *controllerAddress) GetAddressByID(ctx *gin.Context) {
	id_payment := ctx.Param("id_address")
	id_paymentstr, err := strconv.Atoi(id_payment)
	customer, err := c.service.GetAddressByID(id_paymentstr)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Cliente não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)

}

func (c *controllerAddress) UpdateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.BindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdateAddress(address)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseAddress{"Objeto atualizado", "Ok", address}

	ctx.JSON(200, response)

}
