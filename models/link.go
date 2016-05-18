package models

import "github.com/jinzhu/gorm"

type Link struct {
    gorm.Model
    Url string
    Title string
}
