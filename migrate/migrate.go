package main

import (
	"github.com/sayalipol/go-crud/initializers"
	"github.com/sayalipol/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
