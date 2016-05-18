package controllers

import "github.com/jinzhu/gorm"

type BaseController struct {
    db  gorm.DB
}
