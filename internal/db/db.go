package db

import (
	"log"
	"melodyhub/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection to the PostgreSQL database using the provided DSN
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	return db, nil
}

// HandleMigrate runs the database migrations for the Song model
func HandleMigrate(db gorm.DB) error {
	err := db.AutoMigrate(&models.Song{})

	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}
	log.Println("Migration successful")
	return nil
}
