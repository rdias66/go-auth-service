package auth

import (
	"errors"
	jwtService "go-auth-service/services/jwt"
	password "go-auth-service/services/password"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(email, password string) (string, error) // Handles user login
	RefreshToken(token string) (string, error)    // Handles token refresh
}

type authService struct {
	jwtService jwtService.JWTService // Dependency injection of JWTService
	userRepo   UserRepository        // Dependency injection of UserRepository
}

// NewAuthService is a constructor function for AuthService
func NewAuthService(jwtService jwtService.JWTService, userRepo UserRepository) AuthService {
	return &authService{
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

// Login validates the user credentials and returns a JWT token if successful
func (a *authService) Login(email, entryPassword string) (string, error) {
	// Find the user by email
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Verify the password using the password service
	err = password.VerifyPassword(entryPassword, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token if the credentials are valid
	token, err := a.jwtService.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil // Return the generated token
}

// RefreshToken generates a new JWT token from an existing (valid) token
func (a *authService) RefreshToken(token string) (string, error) {
	// Validate the existing token
	validToken, err := a.jwtService.ValidateToken(token)
	if err != nil {
		return "", err
	}

	// Extract the email from the token claims
	claims := validToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	// Find the user by email
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Generate a new JWT token
	newToken, err := a.jwtService.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return newToken, nil // Return the new token
}
