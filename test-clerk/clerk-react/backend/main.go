package main

import (
	"github.com/MC0117/test-clerk/clerk-react/backend/controllers"
	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	//initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}
	router.Use(cors.New(config))

	router.POST("/post", controllers.PostCreate)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/show", controllers.PostShow)
	router.POST("/api/chat", controllers.ChatAction)
	router.Run() // listen and serve on localhost:3000
}
