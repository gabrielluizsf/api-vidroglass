package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mariarobertap/api-vidroglass/controller"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/service"

	//"github.com/mariarobertap/api-vidroglass/middleware"
	//"io"
	//"os"
	"net/http"
)

var (
	clienteService    interfaces.ClienteService    = service.NewClienteService()
	clienteController interfaces.ClienteController = controller.NewClienteController(clienteService)

	PaymentFornService interfaces.PaymentFormService    = service.NewPaymentFormService()
	PaymentController  interfaces.PaymentFormController = controller.NewPaymentFormController(PaymentFornService)

	AddressService    interfaces.AddressService    = service.NewAddressService()
	AddressController interfaces.AddressController = controller.NewAddressController(AddressService)
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

	server.GET("/user/:userid", func(c *gin.Context) {
		userid := c.Param("userid")
		message := "userid is " + userid
		c.String(http.StatusOK, message)
		fmt.Println(message)
	})

	/*
	 [GET] /cliente Retornará clientes
	 [GET] /cliente/id Retornará cliente
	 [POST] /cliente Cadastrar cliente e endereco
	 [GET] /endereco/id Retornará endereco

	 [POST] /tipo_produto Cadastrar tipo_produto
	 [POST] /produto Cadastrar produto
	 [GET] /produto/id Retornará produto

	 [POST] /item Cadastrar item
	 [GET] /item/id Retornará item

	 [POST] /nota Cadastrar nota
	 [GET] /nota/id Retornará nota

	 [GET] /empresa/id Retornará empresa
	*/

	server.Run(":8080")

}
