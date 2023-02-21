package router

import (
	"vidroglass/api/handler"
	"vidroglass/gateway"

	"vidroglass/usecase"

	"github.com/gin-gonic/gin"
)

var (
	customerUseCase gateway.CustomerUseCase        = usecase.NewCustomerUseCase()
	CustomerHandler handler.ClientHandlerInterface = handler.NewClientHandler(customerUseCase)

	productUseCase gateway.ProductUseCase          = usecase.NewProductUseCase()
	ProductHandler handler.ProductHandlerInterface = handler.NewProductHandler(productUseCase)

	paymentUseCase gateway.PaymentUseCase          = usecase.NewPaymentUseCase()
	PaymentHandler handler.PaymentHandlerInterface = handler.NewPaymentHandler(paymentUseCase)

	invoiceUseCase gateway.InvoiceUseCase          = usecase.NewInvoiceUseCase()
	InvoiceHandler handler.InvoiceHandlerInterface = handler.NewInvoiceHandler(invoiceUseCase)
)

func StartRoute(server *gin.Engine) {

	server.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	//	server.Use(gin.Recovery(), middleware.Auth())

	server.GET("/customer", CustomerHandler.GetCustomersHandler)
	server.GET("/customer/:customerID", CustomerHandler.GetCustomerByIDHandler)
	server.POST("/customer", CustomerHandler.CreateCustomerHandler)

	server.GET("/product", ProductHandler.GetProductsHandler)
	server.GET("/product/:productID", ProductHandler.GetProductByIDHandler)
	server.POST("/product", ProductHandler.CreateProductHandler)
	server.GET("/product/type", ProductHandler.GetProductTypesHandler)
	server.POST("/product/type", ProductHandler.CreateProductTypeHandler)

	server.GET("/payment", PaymentHandler.GetPaymentsHandler)
	server.GET("/payment/:paymentID", PaymentHandler.GetPaymentByIDHandler)
	server.POST("/payment", PaymentHandler.CreatePaymentHandler)

	server.GET("/invoice", InvoiceHandler.GetInvoicesHandler)
	server.POST("/invoice", InvoiceHandler.CreateInvoiceHandler)
	server.POST("/invoice/:invoiceID/item", InvoiceHandler.CreateInvoiceItemHandler)

	server.Run(":3000")

}
