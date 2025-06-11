package database

import (
	"log"
	"service-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Book{},
	)

	if err != nil {
		log.Fatalln("❌ Failed to migrate database schema:", err)
	}

	log.Println("✅ Database migration completed")
}
