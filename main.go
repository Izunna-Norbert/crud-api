package main

import (
	"github.com/Izunna-Norbert/crud-api/controllers"
	"github.com/Izunna-Norbert/crud-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
