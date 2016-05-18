package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "github.com/lszanto/links/config"
    "github.com/lszanto/links/models"
)

type LinkController struct {
    db *gorm.DB
    config config.Config
}

// return an instance of controller
func NewLinkController(db *gorm.DB, config config.Config) *LinkController {
    return &LinkController{ db: db, config: config }
}

// CREATE post link
func (lc LinkController) Post(c *gin.Context) {
    // grab parts
    title := c.PostForm("title")
    url := c.PostForm("url")

    // create new
    lc.db.Create(&models.Link{ Title: title, Url: url })

    // created
    c.Status(http.StatusCreated)
}

// READ get link
func (lc LinkController) Get(c *gin.Context) {
    // grab id
    id := c.Params.ByName("id")

    // set link placeholder
    var link models.Link

    // find link
    lc.db.First(&link, id)

    if link.Title == "" {
        c.Status(http.StatusNotFound)
        return
    }

    // return
    c.JSON(http.StatusOK, gin.H{
        "title": link.Title,
        "url": link.Url,
    })
}
