package controllers

import "github.com/jinzhu/gorm"

// BaseController allows controllers to extend off this
type BaseController struct {
    db gorm.DB
}
