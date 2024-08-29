package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	RoleID   string
	Role     Role
}

type Role struct {
	gorm.Model
	Id    string `gorm:"unique;not null"`
	Name  string `gorm:"unique;not null"`
	Users []User
}
