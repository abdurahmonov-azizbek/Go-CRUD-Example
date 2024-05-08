package main

import (
	"github.com/abdurahmonov-azizbek/go-crud/initializers"
	"github.com/abdurahmonov-azizbek/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
