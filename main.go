package main

import (
	"mockinjection/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	ctrl := controller.NewExampleController()
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/", ctrl.ExampleGet)
	router.Run(":8080")
}
