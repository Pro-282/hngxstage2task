package main

import (
	"github.com/pro-282/hngstage2task/initializers"
	"github.com/pro-282/hngstage2task/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
