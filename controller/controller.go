package controller

import (
	"mockinjection/service"

	"github.com/gin-gonic/gin"
)

type exampleController struct{}

func NewExampleController() *exampleController {
	return &exampleController{}
}

func (c exampleController) ExampleGet(ctx *gin.Context) {
	data := service.GetExampleService().Get(1)
	ctx.JSON(200, gin.H{"data": data})
}
