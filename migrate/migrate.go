package main

import (
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/initializers"
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Candidate{})
}
