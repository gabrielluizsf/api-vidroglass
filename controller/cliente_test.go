package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/service"

	"github.com/stretchr/testify/assert"
)

var (
	clienteService interfaces.ClienteService = service.NewClienteService()

	clienteController interfaces.ClienteController = NewClienteController(clienteService)
)

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	router := SetupRouter()
	w := performRequest(router, "GET", "/cliente")
	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)

}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func SetupRouter() *gin.Engine {
	var (
		clienteService    interfaces.ClienteService    = service.NewClienteService()
		clienteController interfaces.ClienteController = NewClienteController(clienteService)
	)

	server := gin.Default()

	server.GET("/cliente", clienteController.FindAll)

	return server

}
