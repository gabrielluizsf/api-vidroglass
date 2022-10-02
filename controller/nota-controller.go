package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerNota struct {
	service interfaces.NotaService
}

func NewNotaController(service interfaces.NotaService) interfaces.NotaController {
	return &controllerNota{
		service: service,
	}
}

func (c *controllerNota) GetNota(ctx *gin.Context) {
	clientes, err := c.service.GetNota()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, clientes)

}

func (c *controllerNota) CreateNota(ctx *gin.Context) {
	var nota models.Nota

	if err := ctx.BindJSON(&nota); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	id, err := c.service.CreateNota(nota)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseNota{"Objeto criado", "Ok", id}

	ctx.JSON(200, response)

}

func (c *controllerNota) GetNotaByID(ctx *gin.Context) {
	id_nota := ctx.Param("id_nota")
	id_notastr, err := strconv.Atoi(id_nota)
	customer, err := c.service.GetNotaByID(id_notastr)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Nota não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)

}

func (c *controllerNota) UpdateNota(ctx *gin.Context) {
	var nota models.Nota

	if err := ctx.BindJSON(&nota); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	nota, err := c.service.UpdateNota(nota)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponseNotaObjetc{"Objeto atualizado", "Ok", nota}

	ctx.JSON(200, response)

}
