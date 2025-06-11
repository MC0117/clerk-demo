package main

import (
	"github.com/MC0117/test-clerk/clerk-react/backend/controllers"
	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.POST("/post", controllers.PostCreate)
	router.Run() // listen and serve on localhost:3000
}
