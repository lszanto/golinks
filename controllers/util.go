package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/lszanto/links/config"
	"github.com/lszanto/links/helpers"
)

// UtilController structure setup
type UtilController struct {
	BaseController
}

// NewUtilController returns a new instance of this controller
func NewUtilController(db *gorm.DB, config config.Config) *UtilController {
	return &UtilController{BaseController{db: db, config: config}}
}

// ImageList function, returns image lister
func (uc UtilController) ImageList(c *gin.Context) {
	// grab images from link
	images, err := helpers.GetImgsFromURL(c.PostForm("url"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error with URL",
		})
		return
	}

	// if no links
	if len(images) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "No images found",
		})
		return
	}

	// return json images if we have enough and url isn't malfunctioning
	c.JSON(http.StatusOK, images)
}
