package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sayalipol/go-crud/controller"
	"github.com/sayalipol/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controller.PostControllers)
	r.GET("/posts", controller.PostsIndex)
	r.GET("/posts/:id", controller.PostsShow)
	r.PUT("/posts/:id", controller.PostUpdate)
	r.DELETE("/posts/:id", controller.PostDelete)

	r.Run()
}
