package models

import (
  "time"
  "github.com/jinzhu/gorm"
)

// User model, defines the user and attributes
type User struct {
    gorm.Model
    Username string
    Password string `json:"-"`
    Email    string `json:"-"`
    CreatedAt *time.Time `json:",omitempty"`
    UpdatedAt *time.Time `json:",omitempty"`
    DeletedAt *time.Time `json:",omitempty"`
}
