package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/lszanto/links/config"
	//"github.com/lszanto/links/models"
)

// UserController structure setup
type UserController struct {
	db     *gorm.DB
	config config.Config
}

// NewUserController returns an instance of controller
func NewUserController(db *gorm.DB, config config.Config) *UserController {
	return &UserController{db: db, config: config}
}

// Login checks password and signs in/returns jwt token
func (uc UserController) Login(c *gin.Context) {
	// check password
	if c.PostForm("password") == "p123" {
		// create token
		token := jwt.New(jwt.SigningMethodHS256)

		// set claims
		token.Claims["foo"] = "bar"
		token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// create token string
		tokenString, err := token.SignedString([]byte(uc.config.SecretKey))

		// check for error
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	} else {
		c.Status(http.StatusUnauthorized)
	}
}
