package auth

import (
	"go-auth-service/models" // Importing the models package

	"gorm.io/gorm" // GORM is our ORM
)

// UserRepository defines the methods for accessing user data
type UserRepository interface {
	FindByEmail(email string) (*models.User, error) // Find a user by their email
}

// userRepository is the struct that implements the UserRepository interface
type userRepository struct {
	db *gorm.DB // A GORM DB instance for interacting with the database
}

// NewUserRepository is a constructor function for UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByEmail fetches a user by their email address
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User // Declare a variable to hold the user data

	// Query the database for the user with the given email
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err // Return an error if the user is not found
	}

	return &user, nil // Return the found user
}
