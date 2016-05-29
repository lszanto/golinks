package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/lszanto/links/config"
	"github.com/lszanto/links/helpers"
)

// SiteController structure setup
type SiteController struct {
	BaseController
}

// NewSiteController returns a new instance of this controller
func NewSiteController(db *gorm.DB, config config.Config) *SiteController {
	return &SiteController{BaseController{db: db, config: config}}
}

// Home function, just returns base template
func (sc SiteController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// ImageList function, returns image lister
func (sc SiteController) ImageList(c *gin.Context) {
	// grab images from website
	images := helpers.GetImgsFromURL("http://weaintgotnohistory.com")

	c.HTML(http.StatusOK, "images.html", gin.H{
		"img_list": images,
	})
}
