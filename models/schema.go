package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"type:uuid;primary_key;unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	RoleID   string
	Role     Role
	gorm.Model
}

type Role struct {
	Id    string `gorm:"type:uuid;primary_key;unique;not null"`
	Name  string `gorm:"unique;not null"`
	Users []User
	gorm.Model
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.New().String()
	return
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.Id = uuid.New().String()
	return
}
