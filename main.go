package main

import (
	"go-auth-service/config"
	"go-auth-service/models"
	"log"
)

func main() {
	configData := config.LoadConfig()

	// Connect to the database
	db, err := config.ConnectDatabase(configData)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Migrate the schema and seed the database
	err = models.MigrateAndSeed(db, configData.AdminEmail, configData.AdminPassword)
	if err != nil {
		log.Fatalf("Could not migrate and seed the database: %v", err)
	}

	log.Println("Auth microservice started successfully.")
}
