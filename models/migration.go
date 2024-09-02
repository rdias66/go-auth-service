package models

import (
	password "go-auth-service/services/password"
	"log"

	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB, adminEmail, adminPassword string) error {
	// Auto-migrate the schema
	err := db.AutoMigrate(&User{}, &Role{})
	if err != nil {
		return err
	}

	// Seed Roles
	roles := []Role{
		{Name: "Admin"},
		{Name: "User"},
	}
	for _, role := range roles {
		// Create roles if they do not exist
		if err := db.FirstOrCreate(&role, Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}

	// Seed Admin User
	var adminRole Role
	if err := db.Where("name = ?", "Admin").First(&adminRole).Error; err != nil {
		return err
	}

	// Hash the admin password using the password service
	hashedPassword, err := password.HashPassword(adminPassword)
	if err != nil {
		return err
	}

	admin := User{
		Email:    adminEmail,
		Password: hashedPassword,
		RoleID:   adminRole.Id, // Use RoleID for foreign key
	}

	if err := db.FirstOrCreate(&admin, User{Email: admin.Email}).Error; err != nil {
		return err
	}

	log.Println("Database migrated and seeded successfully.")
	return nil
}
