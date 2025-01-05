package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCodeSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/doug", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Deveria retornar 200")

	mockResponse := `{"API diz:":"E ai doug, tudo beleza?"}`
	ResponseBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockResponse, string(ResponseBody), "Deveria retornar a mensagem de saudação")
}
