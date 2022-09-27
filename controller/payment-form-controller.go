package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
)

type controllerPaymentForm struct {
	service interfaces.PaymentFormService
}

func NewPaymentFormController(service interfaces.PaymentFormService) interfaces.PaymentFormController {
	return &controllerPaymentForm{
		service: service,
	}
}

func (c *controllerPaymentForm) GetPaymentForm(ctx *gin.Context) {
	clientes, err := c.service.GetPaymentForm()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, clientes)

}

func (c *controllerPaymentForm) CreatePaymentForm(ctx *gin.Context) {
	var payment models.PaymentForm

	if err := ctx.BindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}
	id, err := c.service.CreatePaymentForm(payment)

	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}
	payment.Id_payment = id
	response := models.GoodResponsePayment{"Objeto criado", "Ok", payment}

	ctx.JSON(200, response)

}

func (c *controllerPaymentForm) GetPaymentFormByID(ctx *gin.Context) {
	id_payment := ctx.Param("id_payment")
	id_paymentstr, err := strconv.Atoi(id_payment)
	customer, err := c.service.GetPaymentFormByID(id_paymentstr)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Cliente não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)

}

func (c *controllerPaymentForm) UpdatePaymentForm(ctx *gin.Context) {
	var payment models.PaymentForm

	if err := ctx.BindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, models.BadResponse{"Ocorreu um erro ao criar o Objeto", "Error", err.Error()})
		return
	}

	err := c.service.UpdatePaymentForm(payment)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Ocorreu um erro ao atualizar o objeto", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	response := models.GoodResponsePayment{"Objeto atualizado", "Ok", payment}

	ctx.JSON(200, response)

}

func (c *controllerPaymentForm) DeletePaymentByID(ctx *gin.Context) {
	id_payment := ctx.Param("idpagamento")
	id_paymentstr, err := strconv.Atoi(id_payment)
	if err != nil {
		fmt.Println(err)
		response := models.BadResponse{"Cliente não encontrado", "Error", err.Error()}
		ctx.JSON(400, response)
		return
	}

	err = c.service.DeletePaymentByID(id_paymentstr)

	if err != nil {
		response := models.BadResponse{
			Message: "Erro ao deletar o objeto",
			Status:  "Erro",
			Erro:    err.Error()}

		ctx.JSON(500, response)
		return
	}
	response := models.GoodResponsePayment{
		Message:     "Objeto Deletado",
		Status:      "Ok",
		PaymentForm: models.PaymentForm{}}

	ctx.JSON(200, response)

}
