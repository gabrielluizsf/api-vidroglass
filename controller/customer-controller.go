package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/models"
    "github.com/mariarobertap/api-vidroglass/service"

)



type CustomerController interface {
	FindAll(ctx *gin.Context) 
	Save(ctx *gin.Context) models.Customer
}

type controller struct {
	service service.CustomerService
}


func NewCustomerController(service service.CustomerService) CustomerController {
	return &controller {
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context){
	customers, err := c.service.FindAll()
	if(err != nil){
		fmt.Println(err)
		ctx.JSON(400, "error")
	}

	ctx.JSON(200, customers)


}


func (c *controller) Save(ctx *gin.Context) models.Customer{
	var customer models.Customer
	ctx.BindJSON(&customer)
	c.service.Save(customer)

	return customer
}