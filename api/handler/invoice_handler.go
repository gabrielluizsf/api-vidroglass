package handler

import (
	"fmt"
	"net/http"
	"vidroglass/api/util"
	"vidroglass/gateway"
	"vidroglass/model"

	"github.com/gin-gonic/gin"
)

type InvoiceHandlerInterface interface {
	GetInvoicesHandler(ctx *gin.Context)
	CreateInvoiceHandler(ctx *gin.Context)
	CreateInvoiceItemHandler(ctx *gin.Context)
}

type invoiceHandler struct {
	invoiceUseCase gateway.InvoiceUseCase
}

func NewInvoiceHandler(usecase gateway.InvoiceUseCase) InvoiceHandlerInterface {
	return &invoiceHandler{
		invoiceUseCase: usecase,
	}
}

func (c *invoiceHandler) GetInvoicesHandler(ctx *gin.Context) {

	response, err := c.invoiceUseCase.GetInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "response")
	}
	ctx.JSON(200, response)
}

func (c *invoiceHandler) CreateInvoiceHandler(ctx *gin.Context) {
	var invoice model.Invoice

	if err := ctx.BindJSON(&invoice); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	invoice, err := c.invoiceUseCase.CreateInvoice(invoice)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  invoice,
	}

	ctx.JSON(200, response)
}

func (c *invoiceHandler) CreateInvoiceItemHandler(ctx *gin.Context) {
	var item model.Item
	invoiceId, err := getInvoiceID(ctx)
	if err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "invoice_id inválido!")
		ctx.JSON(http.StatusBadRequest, response)
	}

	if err := ctx.BindJSON(&item); err != nil {
		response := util.BuildErrorResponse("A requisição está incorreta!", "Erro na requisição")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	item.InvoiceID = invoiceId
	item, err = c.invoiceUseCase.CreateInvoiceItem(item)

	if err != nil {
		fmt.Println(err)
		response := util.BuildErrorResponse("Ocorreu um erro ao criar o Objeto", "Error")
		ctx.JSON(500, response)
		return
	}

	response := model.Reponse{
		Type:    "Objeto criado.",
		Success: true,
		Detail:  item,
	}

	ctx.JSON(200, response)
}
