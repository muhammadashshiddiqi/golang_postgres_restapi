package main

import (
	"go-crud-postgresql/initializers"
	model "go-crud-postgresql/models"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnecToDB()
}

func main()  {
	initializers.DB.AutoMigrate(&model.Post{})
}