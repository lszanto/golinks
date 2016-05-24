package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model, defines the user and attributes
type User struct {
	gorm.Model
	Username  string     `json:"username" form:"username" binding:"required" sql:"unique"`
	Password  string     `json:"-" form:"passowrd" binding:"required" sql:"size:60"`
	Email     string     `json:"-"`
	CreatedAt *time.Time `json:",omitempty"`
	UpdatedAt *time.Time `json:",omitempty"`
	DeletedAt *time.Time `json:",omitempty"`
	Links     []Link     `json:",omitempty"`
}
