package handler

import (
	"fmt"
	"net/http"
	"vidroglass/api/util"
	"vidroglass/gateway"

	"vidroglass/model"

	"github.com/gin-gonic/gin"
)

type ClientHandlerInterface interface {
	GetCustomersHandler(ctx *gin.Context)
	CreateCustomerHandler(ctx *gin.Context)
	GetCustomerByIDHandler(ctx *gin.Context)
}

type customerHandler struct {
	customerUseCase gateway.CustomerUseCase
}

func NewClientHandler(usecase gateway.CustomerUseCase) ClientHandlerInterface {
	return &customerHandler{
		customerUseCase: usecase,
	}
}

func (c *customerHandler) GetCustomersHandler(ctx *gin.Context) {

	response, err := c.customerUseCase.GetCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "response")

	}
	ctx.JSON(200, response)
}

func (c *customerHandler) CreateCustomerHandler(ctx *gin.Context) {
	var customer model.Customer

	if err := ctx.BindJSON(&customer); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	customer, err := c.customerUseCase.CreateCustomer(customer)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  customer,
	}

	ctx.JSON(200, response)
}

func (c *customerHandler) GetCustomerByIDHandler(ctx *gin.Context) {

	customerId, err := getCustomerID(ctx)
	if err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "customer_id inválido!")
		ctx.JSON(http.StatusBadRequest, response)
	}

	customer, err := c.customerUseCase.GetCustomerByID(customerId)

	if err != nil {
		response := util.BuildErrorResponse("Customer not found!", "not found")
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, customer)
}
