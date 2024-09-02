package controllers

import (
	"encoding/json"
	"go-auth-service/services/auth"
	"go-auth-service/utils"
	"net/http"
)

// Dependency Injection
var authService auth.AuthService

// Ininitializes the controller with the provided AuthService
func InitializeAuthController(service auth.AuthService) {
	authService = service
}

// Handles user login and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	token, err := authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// Handles refreshing the JWT token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Extract the token from request headers
	token := r.Header.Get("Authorization")

	if token == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Token is required")
		return
	}

	newToken, err := authService.RefreshToken(token)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": newToken})
}

// Handles user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	// Implement logout logic, such as invalidating the token or session
	// This example assumes a stateless approach, so this may be a no-op

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Logged out"})
}
