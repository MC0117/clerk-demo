package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/MC0117/test-clerk/clerk-react/backend/models"
)

func PostCreate(c *gin.Context) {

	post := models.Post{Title: "tajtle", Body: "myBody"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
