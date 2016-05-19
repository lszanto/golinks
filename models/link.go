package models

import (
  "time"
  "github.com/jinzhu/gorm"
)

// Link structure, defines the link model
type Link struct {
    gorm.Model
    URL   string
    Title string
    DeletedAt *time.Time `json:",omitempty"`
    User *User `json:",omitempty"`
    UserID int `json:"-"`
}
