package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"mockinjection/controller"
	"mockinjection/mocks"
	"mockinjection/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupServer() *gin.Engine {
	ctrl := controller.NewExampleController()
	router := gin.Default()
	router.GET("/", ctrl.ExampleGet)
	return router
}

type Result struct {
	Data int `json:"data"`
}

func TestInterface(t *testing.T) {
	testRouter := SetupServer()
	mockService := &mocks.ExampleService{}
	// Get service return 값 11로 변경.
	mockService.On("Get", 1).Return(11)
	service.SetExampleService(mockService)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)

	var result Result
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyByte, &result)
	// 실 서비스에선 1을 return하는데, test에서는 11을 return함.
	assert.Equal(t, 11, result.Data)
}
