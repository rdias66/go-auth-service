package main

import (
	"go-auth-service/config"
	"go-auth-service/controllers"
	"go-auth-service/models"
	"go-auth-service/routes"
	auth "go-auth-service/services/auth"
	jwtService "go-auth-service/services/jwt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from env
	configData := config.LoadConfig()

	// Connect to the database
	db, err := config.ConnectDatabase(configData)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Migrate the schema and seed the database with an admin user
	err = models.MigrateAndSeed(db, configData.AdminEmail, configData.AdminPassword)
	if err != nil {
		log.Fatalf("Could not migrate and seed the database: %v", err)
	}

	// Initialize repositories(user service)
	userRepo := auth.NewUserRepository(db)

	// Initialize JWT service
	jwtServiceInstance := jwtService.NewJWTService(configData.JWTSecret, "AuthMicroservice")

	// Initialize authentication service
	authInstance := auth.NewAuthService(jwtServiceInstance, userRepo)

	// Initialize controllers
	controllers.InitializeAuthController(authInstance)

	// Set up HTTP router
	router := mux.NewRouter()
	routes.RegisterAuthRoutes(router)

	// Start the HTTP server
	log.Println("Auth microservice started successfully.")
	if err := http.ListenAndServe(":"+configData.ServerPort, router); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
