package main

import (
	"github.com/brothermos/go-crud/initializers"
	model "github.com/brothermos/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
