package main

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/lszanto/links/config"
)

type Link struct {
    gorm.Model
    Url string
    Title string
}

func JWTAuth(key string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // parse jwt
        token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
            return []byte(key), nil
        })

        if err == nil && token.Valid {
            // move on
            c.Next()
        } else {
            // abort
            c.AbortWithStatus(http.StatusUnauthorized)
        }
    }
}

func main() {
    // set error holder
    var err error

    // open a config file
    config_file, _ := os.Open("app/config.json")

    // decode config
    decoder := json.NewDecoder(config_file)

    // create config object
    config := config.Config{}

    // decode
    err = decoder.Decode(&config)

    // create db connection
    db, err := gorm.Open(config.DB_Engine, config.DB_String)

    if err != nil {
        panic("failed to connect to database")
    }

    // setup router
    router := gin.Default()

    // add routes
    router.POST("/login", func(c *gin.Context) {
        // check password
        if c.PostForm("password") == "p123" {
            // create token
            token := jwt.New(jwt.SigningMethodHS256)

            // set claims
            token.Claims["foo"] = "bar"
            token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

            // create token string
            tokenString, err := token.SignedString([]byte(config.SESS_Secret))

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
    })

    router.POST("/link", JWTAuth(config.SESS_Secret), func(c *gin.Context) {
        // grab parts
        title := c.PostForm("title")
        url := c.PostForm("url")

        // create new
        db.Create(&Link{ Title: title, Url: url })

        // created
        c.Status(http.StatusCreated)
    })

    router.GET("/link/:id", func(c *gin.Context) {
        // grab id
        id := c.Params.ByName("id")

        // set link placeholder
        var link Link

        // find link
        db.First(&link, id)

        if link.Title == "" {
            c.Status(http.StatusNotFound)
            return
        }

        // return
        c.JSON(http.StatusOK, gin.H{
            "title": link.Title,
            "url": link.Url,
        })
    })

    // setup static folder
    router.Static("/assets", "./assets")

    // lets go
    fmt.Println("Lets do this")

    // run server
    router.Run()
}
