package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("❌ Failed to load the credentials from the environment file:", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalln("❌ DATABASE_URL not found in environment variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("❌ Failed to connect to the database:", err)
	}

	log.Println("✅ Successfully connected to the database")
	return db
}
