package models

import "github.com/jinzhu/gorm"

// User model, defines the user and attributes
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}
