package main

import (
	"go-crud-v1/initalizers"
	"go-crud-v1/models"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
