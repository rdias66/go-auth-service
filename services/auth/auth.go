package auth

import (
	"errors"
	jwtService "go-auth-service/services/jwt"
	password "go-auth-service/services/password"

	"fmt"

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
func (a *authService) Login(email string, entryPassword string) (string, error) {

	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = password.VerifyPassword(entryPassword, user.Password)
	fmt.Println(err)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := a.jwtService.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Generates a new JWT token from an existing (valid) token
func (a *authService) RefreshToken(token string) (string, error) {

	validToken, err := a.jwtService.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims := validToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	newToken, err := a.jwtService.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
