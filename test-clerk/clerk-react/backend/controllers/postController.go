package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MC0117/test-clerk/clerk-react/backend/clerk"
	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/MC0117/test-clerk/clerk-react/backend/models"
)

var body struct {
	Title string
	Body  string
}

func ChatAction(c *gin.Context) {
	rawAuth := c.GetHeader("Authorization")
	fmt.Println("Header received:", rawAuth)

	if rawAuth == "" || !strings.HasPrefix("Bearer ") {
		c.JSON(401, gin.H{"error": "Missing or Invalid authorization header"})
		return
	}

	token := strings.TrimPrefix(rawAuth, "Bearer ")

	clerk.GetUserId(c, token)
	var chat_req struct {
		Message string `json:"message"`
		Agent   string `json:"agent"`
	}

	fmt.Println("Post req received...")
	err := c.Bind(&chat_req)
	if err != nil {
		fmt.Println("Error binding request:", err)
		return
	}
	fmt.Println("Bind successful")
	fmt.Println("Message:", chat_req.Message)
	fmt.Println("Agent:", chat_req.Agent)

}

func PostCreate(c *gin.Context) {
	//request body

	//not sure what this does
	c.Bind(&body)

	//Prepare query
	post := models.Post{Title: body.Title, Body: body.Body}

	//execute
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
	//post := models.Post{Title: body.Title, Body: body.Body}

	//result := initializers.DB.Where("ID = ?", 0).Find(&post)

}

func PostShow(c *gin.Context) {

	var post models.Post

	initializers.DB.Find(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
	//post := models.Post{Title: body.Title, Body: body.Body}

	//result := initializers.DB.Where("ID = ?", 0).Find(&post)

}
