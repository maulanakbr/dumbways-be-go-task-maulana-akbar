package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/controllers"
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/candidates", controllers.CreateCandidate)
	r.PUT("/candidates/:id", controllers.UpdateCandidateByPk)
	r.DELETE("/candidates/:id", controllers.DeleteCandidateByPk)
	r.GET("/candidates", controllers.FindAllCandidates)
	r.GET("/candidates/:id", controllers.FindCandidateByPk)

	r.Run()
}
