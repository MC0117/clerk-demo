package controllers

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"

	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/MC0117/test-clerk/clerk-react/backend/models"

	"github.com/golang-jwt/jwt/v4"
)

var body struct {
	Title string
	Body  string
}

func ChatAction(c *gin.Context) {
	// Extract and validate Authorization header
	rawAuth := c.GetHeader("Authorization")
	if rawAuth == "" || !strings.HasPrefix(rawAuth, "Bearer ") {
		c.JSON(401, gin.H{"error": "Missing or Invalid authorization header"})
		return
	}

	// Extract JWT token from Bearer string
	token_str := strings.TrimPrefix(rawAuth, "Bearer ")

	// Get JWKS URL from environment
	jwksURL := os.Getenv("CLERK_JWKS_URL")
	if jwksURL == "" {
		c.JSON(500, gin.H{"error": "JWKS URL not configured"})
		return
	}

	// Initialize JWKS client with refresh options
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			log.Printf("JWKS refresh error: %s", err)
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to initialize JWKS client"})
		return
	}

	// Parse and validate JWT token
	token, err := jwt.Parse(token_str, jwks.Keyfunc)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	if !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	// Parse chat request body
	var chat_req struct {
		Message string `json:"message"`
		Agent   string `json:"agent"`
	}

	if err := c.Bind(&chat_req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Log request details
	log.Printf("Chat request - Message: %s, Agent: %s", chat_req.Message, chat_req.Agent)

	// TODO: Process chat request
	c.JSON(200, gin.H{"status": "success"})
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
