package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/dgrijalva/jwt-go"

    "github.com/lszanto/links/config"
    //"github.com/lszanto/links/models"
)

type UserController struct {
    db *gorm.DB
    config config.Config
}

// return an instance of controller
func NewUserController(db *gorm.DB, config config.Config) *UserController {
    return &UserController{ db: db, config: config }
}

// login
func (uc UserController) Login(c *gin.Context) {
    // check password
    if c.PostForm("password") == "p123" {
        // create token
        token := jwt.New(jwt.SigningMethodHS256)

        // set claims
        token.Claims["foo"] = "bar"
        token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

        // create token string
        tokenString, err := token.SignedString([]byte(uc.config.SESS_Secret))

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
