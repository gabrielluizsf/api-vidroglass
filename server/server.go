package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mariarobertap/api-vidroglass/controller"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/service"
	//"github.com/mariarobertap/api-vidroglass/middleware"
	//"io"
	//"os"
)

var (
	clienteService    interfaces.ClienteService    = service.NewClienteService()
	clienteController interfaces.ClienteController = controller.NewClienteController(clienteService)

	PaymentFornService interfaces.PaymentFormService    = service.NewPaymentFormService()
	PaymentController  interfaces.PaymentFormController = controller.NewPaymentFormController(PaymentFornService)

	AddressService    interfaces.AddressService    = service.NewAddressService()
	AddressController interfaces.AddressController = controller.NewAddressController(AddressService)

	ProductTypeService    interfaces.ProductTypeService    = service.NewProductTypeService()
	ProductTypeController interfaces.ProductTypeController = controller.NewProductTypeController(ProductTypeService)

	ProductService    interfaces.ProductService    = service.NewProductService()
	ProductController interfaces.ProductController = controller.NewProductController(ProductService)
)

/*
func setupLogOutput(){
	f, err := os.Create("gin.log")

	if(err != nil){
		fmt.Println(err)
		return
	}

	gin.DefaultWriter =  io.MultiWriter(f, os.Stdout)
}
*/

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//setupLogOutput()
	server := gin.Default()

	//server.Use(gin.Recovery(), middleware.Logger(), middleware.Auth())

	server.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.GET("/cliente", clienteController.FindAll)
	server.POST("/cliente", clienteController.Save)
	server.GET("/cliente/:id_cliente", clienteController.GetClientById)
	server.PUT("/cliente", clienteController.UpdateClientById)

	server.GET("/paymentform", PaymentController.GetPaymentForm)
	server.POST("/paymentform", PaymentController.CreatePaymentForm)
	server.GET("/paymentform/:id_payment", PaymentController.GetPaymentFormByID)
	server.PUT("/paymentform", PaymentController.UpdatePaymentForm)

	server.GET("/address", AddressController.GetAddress)
	server.POST("/address", AddressController.CreateAddress)
	server.GET("/address/:id_address", AddressController.GetAddressByID)
	server.PUT("/address", AddressController.UpdateAddress)

	server.GET("/product/type", ProductTypeController.GetProductType)
	server.POST("/product/type", ProductTypeController.CreateProductType)
	server.GET("/product/type/:id_tipo", ProductTypeController.GetProductTypeByID)
	server.PUT("/product/type", ProductTypeController.UpdateProductType)

	server.GET("/product", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id_tipo", ProductController.GetProductByID)
	server.PUT("/product", ProductController.UpdateProduct)

	server.Run(":8080")

}
