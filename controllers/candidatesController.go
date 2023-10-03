package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/initializers"
	"github.com/maulanakbr/dumbways-be-go-task-maulana-akbar/models"
)

func CreateCandidate(c *gin.Context) {
	var body struct {
		Name   string
		Vision string
	}

	c.Bind(&body)

	candidate := models.Candidate{Name: body.Name, Vision: body.Vision}
	result := initializers.DB.Create(&candidate)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"candidate": candidate,
	})
}

func FindAllCandidates(c *gin.Context) {
	var candidates []models.Candidate
	result := initializers.DB.Find(&candidates)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"candidates": candidates,
	})
}

func FindCandidateByPk(c *gin.Context) {
	id := c.Param("id")

	var candidate models.Candidate
	result := initializers.DB.First(&candidate, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"candidates": candidate,
	})
}

func UpdateCandidateByPk(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name   string
		Vision string
	}

	c.Bind(&body)

	var candidate models.Candidate
	findCandidate := initializers.DB.First(&candidate, id)

	if findCandidate.Error != nil {
		c.Status(400)
		return
	}

	initializers.DB.Model(&candidate).Updates(models.Candidate{
		Name:   body.Name,
		Vision: body.Vision,
	})

	c.JSON(200, gin.H{
		"candidates": candidate,
	})
}

func DeleteCandidateByPk(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Candidate{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted candidate successfully",
	})
}
