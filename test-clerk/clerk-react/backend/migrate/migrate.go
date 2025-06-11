package migrate

import (
	"github.com/MC0117/test-clerk/clerk-react/backend/initializers"
	"github.com/MC0117/test-clerk/clerk-react/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
