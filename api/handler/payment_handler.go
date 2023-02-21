package handler

import (
	"fmt"
	"net/http"
	"vidroglass/api/util"
	"vidroglass/gateway"

	"vidroglass/model"

	"github.com/gin-gonic/gin"
)

type PaymentHandlerInterface interface {
	GetPaymentsHandler(ctx *gin.Context)
	CreatePaymentHandler(ctx *gin.Context)
	GetPaymentByIDHandler(ctx *gin.Context)
}

type paymentHandler struct {
	paymentUseCase gateway.PaymentUseCase
}

func NewPaymentHandler(usecase gateway.PaymentUseCase) PaymentHandlerInterface {
	return &paymentHandler{
		paymentUseCase: usecase,
	}
}

func (c *paymentHandler) GetPaymentsHandler(ctx *gin.Context) {

	response, err := c.paymentUseCase.GetPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "response")

	}
	ctx.JSON(200, response)
}

func (c *paymentHandler) CreatePaymentHandler(ctx *gin.Context) {
	var payment model.PaymentMethod

	if err := ctx.BindJSON(&payment); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	payment, err := c.paymentUseCase.CreatePayment(payment)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  payment,
	}

	ctx.JSON(200, response)
}

func (c *paymentHandler) GetPaymentByIDHandler(ctx *gin.Context) {

	paymentId, err := getPaymentID(ctx)
	if err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "payment_id inválido!")
		ctx.JSON(http.StatusBadRequest, response)
	}

	payment, err := c.paymentUseCase.GetPaymentByID(paymentId)

	if err != nil {
		response := util.BuildErrorResponse("Payment not found!", "not found")
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, payment)
}
