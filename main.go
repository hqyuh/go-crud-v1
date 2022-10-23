package main

import (
	"go-crud-v1/initalizers"
	"go-crud-v1/routes"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDB()
}

func main() {
	routes.SetupRoutes()
}
