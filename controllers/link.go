package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "github.com/lszanto/links/config"
    "github.com/lszanto/links/models"
)

// LinkController structure setup
type LinkController struct {
    db     *gorm.DB
    config config.Config
}

// NewLinkController returns a new instance of this controller
func NewLinkController(db *gorm.DB, config config.Config) *LinkController {
    return &LinkController{db: db, config: config}
}

// Create func to create a new link
func (lc LinkController) Create(c *gin.Context) {
    // grab parts
    title := c.PostForm("title")
    url := c.PostForm("url")

    // create new
    lc.db.Create(&models.Link{Title: title, URL: url})

    // created
    c.JSON(http.StatusCreated, gin.H{
        "status": http.StatusCreated,
        "message": "Link created",
    })
}

// Get a singular link via id
func (lc LinkController) Get(c *gin.Context) {
    // grab id
    id := c.Params.ByName("id")

    // set link placeholder
    var link models.Link

    // find link
    lc.db.First(&link, id)

    if link.Title == "" {
        c.JSON(http.StatusNotFound, gin.H{
            "status": http.StatusNotFound,
            "message": "Link not found",
        })
        return
    }

    // return
    c.JSON(http.StatusOK, link)
}

// Update a link
func (lc LinkController) Update(c *gin.Context) {
    // grab id
    id := c.Params.ByName("id")

    // set link placeholder
    var link models.Link

    // find link
    lc.db.First(&link, id)

    if link.Title == "" {
        c.JSON(http.StatusNotFound, gin.H{
            "status": http.StatusNotFound,
            "message": "Link not found",
        })
        return
    }

    // update fields
    lc.db.Model(&link).Updates(models.Link{ Title: c.PostForm("title"), URL: c.PostForm("url") })

    // send accepted response
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "message": "Link updated",
    })
}

// GetAll links
func (lc LinkController) GetAll(c *gin.Context) {
    // set links placeholder
    var links[] models.Link

    // grab all links
    lc.db.Find(&links)

    // return all links
    c.JSON(http.StatusOK, gin.H{
        "links": links,
    })
}

// Delete a link
func (lc LinkController) Delete(c *gin.Context) {
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

    // delete link
    lc.db.Delete(&link)

    // return status code
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "message": "Link deleted",
    })
}
