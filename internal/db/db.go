package db

import (
	"gorm.io/gorm"
	"log"
	"melodyhub/pkg/models"
)

func HandleMigrate(db gorm.DB) error {
	err := db.AutoMigrate(&models.Song{})

	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}
	log.Println("Migration successful")
	return nil
}
