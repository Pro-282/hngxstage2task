package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pro-282/hngstage2task/controllers"
	"github.com/pro-282/hngstage2task/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/api", controllers.PostCreate)
	r.GET("/api", controllers.FetchAllUsers)
	r.GET("/api/:id", controllers.FetchUserById)
	r.PUT("/api/:id", controllers.UpdateUserName)
	r.DELETE("/api/:id", controllers.UserDelete)

	r.Run()
}
