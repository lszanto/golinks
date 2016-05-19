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
    return &UserController{ BaseController { db: db, config: config }}
}

// Login checks password and signs in/returns jwt token
func (uc UserController) Login(c *gin.Context) {
    // attempt login
    if uc.login(c.PostForm("username"), c.PostForm("password")) {
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
            "password": uc.hash("p123"),
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
    uc.db.Create(&models.User{ Username: username, Password: password, Email: email })

    // return success
    c.JSON(http.StatusCreated, gin.H{
        "status": http.StatusCreated,
        "message": "User created",
    })
}

// login, attempts to login a user
func (uc UserController) login(username string, password string) bool {
    // create user holder
    var user models.User

    // attempt to find user
    uc.db.Where("username = ?", username).First(&user)

    if user.Username == "" {
        return false
    }

    // check if password matches
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

    // check if we passed(nil = match)
    if err == nil {
        return true
    }

    // if we get to here we've failed
    return false
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
