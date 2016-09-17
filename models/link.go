package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Link structure, defines the link model
type Link struct {
	gorm.Model
	URL       string     `json:"url" form:"url" binding:"required"`
	Title     string     `json:"title" form:"title" binding:"required"`
	DeletedAt *time.Time `json:",omitempty"`
	User      *User      `json:",omitempty"`
	UserID    int        `json:"-"`
}
