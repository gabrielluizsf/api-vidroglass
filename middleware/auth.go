package middleware

import (
	"github.com/gin-gonic/gin"
)

//TODO - Add jwt token
func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"maria": "senha"})
}
