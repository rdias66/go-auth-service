package models

import (
	password "go-auth-service/service"
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
	roles := []string{"Admin", "User"}
	for _, roleName := range roles {
		var role Role
		if err := db.FirstOrCreate(&role, Role{Name: roleName}).Error; err != nil {
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
		Role:     adminRole,
	}

	if err := db.FirstOrCreate(&admin, User{Email: admin.Email}).Error; err != nil {
		return err
	}

	log.Println("Database migrated and seeded successfully.")
	return nil
}
