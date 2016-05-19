package models

import "github.com/jinzhu/gorm"

// Link structure, defines the link model
type Link struct {
	gorm.Model
	URL   string
	Title string
}
