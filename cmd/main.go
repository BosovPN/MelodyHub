package main

import (
	"fmt"
	"log"
	"melodyhub/internal/config"
	"melodyhub/internal/db"
	"melodyhub/internal/handler"
	"melodyhub/internal/repository"
	"melodyhub/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := db.HandleMigrate(*database); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Set up repository and service
	repos := repository.NewRepository(database)
	services := service.NewService(repos)

	// Set up handler with routes
	r := mux.NewRouter()
	handler.RegisterRoutes(r, services)

	port := ":" + cfg.Port

	log.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
