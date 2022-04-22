package interfaces



import (
	"github.com/mariarobertap/api-vidroglass/models"
	"github.com/gin-gonic/gin"
)
type ClienteController interface {
	FindAll(ctx *gin.Context) 
	Save(ctx *gin.Context) 
	UpdateClientById(ctx *gin.Context) 
	GetClientById(ctx *gin.Context)
	
}

type ClienteService interface {


	GetClientById(id_cliente int) (models.Cliente, error)
	FindAll() ([]models.Cliente, error)
	Save(models.Cliente) (int, error)
	UpdateClientById(models.Cliente) (error)


}