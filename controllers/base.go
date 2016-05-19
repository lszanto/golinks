package controllers

import (
    "github.com/jinzhu/gorm"
    "github.com/lszanto/links/config"
)

// BaseController allows controllers to extend off this
type BaseController struct {
    db *gorm.DB
    config config.Config
}
