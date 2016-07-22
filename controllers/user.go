package controllers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/lszanto/links/config"
	"github.com/lszanto/links/models"
)

// UserController structure setup
type UserController struct {
	BaseController
}

// NewUserController returns an instance of controller
func NewUserController(db *gorm.DB, config config.Config) *UserController {
	return &UserController{BaseController{db: db, config: config}}
}

// Login checks password and signs in/returns jwt token
func (uc UserController) Login(c *gin.Context) {
	// attempt login
	user, err := uc.login(c.PostForm("username"), c.PostForm("password"))

	// attempt login
	if err != true {
		// create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid":      user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

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

// CreateUser creates a new user account
func (uc UserController) CreateUser(c *gin.Context) {
	// grab sent attributes
	username := c.PostForm("username")
	password := uc.hash(c.PostForm("password"))
	email := c.PostForm("email")

	// insert user
	uc.db.Create(&models.User{Username: username, Password: password, Email: email})

	// return success
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created",
	})
}

// Get grabs user details via id
func (uc UserController) Get(c *gin.Context) {
	// grab id
	id := c.Params.ByName("id")

	// set user placeholder
	var user models.User

	// find user details
	uc.db.Preload("Links").First(&user, id)

	if user.Username == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	// return
	c.JSON(http.StatusOK, user)
}

// login, attempts to login a user
func (uc UserController) login(username string, password string) (models.User, bool) {
	// create user holder
	var user models.User

	// attempt to find user
	uc.db.Where("username = ?", username).First(&user)

	if user.Username == "" {
		return user, true
	}

	// check if password matches
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// check if we passed(nil = match)
	if err == nil {
		return user, false
	}

	// if we get to here we've failed
	return user, true
}

// hash, returns a hashed password
func (uc UserController) hash(hashString string) string {
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(hashString), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	// return hashed password
	return string(hash)
}
