package main

import (
	"fmt"
	"log"
	"melodyhub/internal/config"
	"melodyhub/internal/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("main.go runnging")
	cfg := config.LoadConfig()

	// Connect to the database
	log.Println("main.go Connect to the database")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
                   cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Run migrations
	log.Println("main.go Run migrations")
	db.HandleMigrate(*database)

	log.Println("main.go ending")
}
