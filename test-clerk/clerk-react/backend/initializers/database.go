package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/MC0117/test-clerk/clerk-react/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	fmt.Printf("Attempting to connect to database with DSN: %s\n", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	fmt.Println("Successfully connected to database!")

	// Auto migrate the database
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
