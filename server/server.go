package main

import (
	"log"

	"github.com/gin-contrib/cors"
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

	ItemService    interfaces.ItemService    = service.NewItemService()
	ItemController interfaces.ItemController = controller.NewItemController(ItemService)

	NotaService    interfaces.NotaService    = service.NewNotaService()
	NotaController interfaces.NotaController = controller.NewNotaController(NotaService)
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

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

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
	server.DELETE("/cliente/:id_cliente", clienteController.DeleteClientById)

	server.GET("/paymentform", PaymentController.GetPaymentForm)
	server.POST("/paymentform", PaymentController.CreatePaymentForm)
	server.GET("/paymentform/:id_payment", PaymentController.GetPaymentFormByID)
	server.PUT("/paymentform", PaymentController.UpdatePaymentForm)
	server.DELETE("/paymentform/:idpagamento", PaymentController.DeletePaymentByID)

	server.GET("/address", AddressController.GetAddress)
	server.POST("/address", AddressController.CreateAddress)
	server.GET("/address/:id_address", AddressController.GetAddressByID)
	server.PUT("/address", AddressController.UpdateAddress)
	server.DELETE("/address/:id_address", AddressController.DeleteAddressByID)

	server.GET("/product/type", ProductTypeController.GetProductType)
	server.POST("/product/type", ProductTypeController.CreateProductType)
	server.GET("/product/type/:id_tipo", ProductTypeController.GetProductTypeByID)
	server.PUT("/product/type", ProductTypeController.UpdateProductType)
	server.DELETE("/product/type/:id_tipo", ProductTypeController.DeleteProductTypeByID)

	server.GET("/product", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id_tipo", ProductController.GetProductByID)
	server.PUT("/product", ProductController.UpdateProduct)
	server.DELETE("/product/:id_product", ProductController.DeleteProductById)

	server.GET("/item", ItemController.FindAll)
	server.POST("/item", ItemController.Save)
	server.GET("/item/:id_item", ItemController.GetItemById)
	server.PUT("/item", ItemController.UpdateItemById)

	server.GET("/nota", NotaController.GetNota)
	server.POST("/nota", NotaController.CreateNota)
	server.GET("/nota/:id_nota", NotaController.GetNotaByID)
	server.PUT("/nota", NotaController.UpdateNota)

	server.Run(":3000")

}
