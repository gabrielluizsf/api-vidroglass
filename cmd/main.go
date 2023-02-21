package main

import (
	"vidroglass/api/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//	+ := godotenv.Load("../.env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	server := gin.Default()

	enableCors(server)
	router.StartRoute(server)

}

func enableCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
}
