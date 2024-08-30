package jwt

import (
	"time" // For handling token expiration time

	"github.com/golang-jwt/jwt/v5" // Importing the JWT package
)

// JWTService is an interface that defines the methods our JWT service will have.
type JWTService interface {
	GenerateToken(userID string, email string) (string, error) // Generate a token
	ValidateToken(token string) (*jwt.Token, error)            // Validate a token
}

// jwtService is a struct that implements JWTService interface. It holds configuration like the secret key and the issuer.
type jwtService struct {
	secretKey string // The key used to sign the token
	issuer    string // The issuer of the token (usually your app's name)
}

// NewJWTService is a constructor function that creates a new jwtService with a given secret key.
func NewJWTService(secretKey string, issuerName string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuerName, // Replace with your actual service name, usually the app name
	}
}

// GenerateToken generates a new JWT token for a given userID and email.
func (j *jwtService) GenerateToken(userID string, email string) (string, error) {
	// Creating a map of claims. Claims are the information stored in the token.
	claims := &jwt.MapClaims{
		"iss":   j.issuer,                              // Issuer claim
		"sub":   userID,                                // Subject claim (typically user ID)
		"email": email,                                 // Email claim
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Expiration time claim (72 hours in this case)
	}

	// Create a new token object with the specified signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(j.secretKey))
}

// ValidateToken validates a given JWT token string and returns the parsed token if it's valid.
func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse the token using the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil // Returning the secret key for validation
	})

	if err != nil {
		return nil, err // Return an error if the token is invalid
	}

	return token, nil // Return the valid token
}
