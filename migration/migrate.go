package main

import (
	"github.com/Izunna-Norbert/crud-api/initializers"
	"github.com/Izunna-Norbert/crud-api/models"
)


func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main () {
	initializers.DB.AutoMigrate(&models.Post{})
}