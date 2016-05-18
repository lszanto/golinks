package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/lszanto/links/config"
    "github.com/lszanto/links/controllers"
)

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

    // CONFIG

    // open a config file
    config_file, _ := os.Open("app/config.json")

    // decode config
    decoder := json.NewDecoder(config_file)

    // create config object
    config := config.Config{}

    // decode
    err = decoder.Decode(&config)

    // GORM DATABASE

    // create db connection
    db, err := gorm.Open(config.DB_Engine, config.DB_String)

    if err != nil {
        panic("failed to connect to database")
    }

    // controllers
    lc := controllers.NewLinkController(db, config)
    uc := controllers.NewUserController(db, config)

    // ROUTER

    // setup router
    router := gin.Default()

    // add routes
    router.POST("/login", uc.Login)
    router.POST("/link", JWTAuth(config.SESS_Secret), lc.Post)
    router.GET("/link/:id", lc.Get)

    // lets go
    fmt.Println("Lets do this")

    // SET STATIC DIR, START SERVER

    // setup static folder
    router.Static("/assets", "./assets")

    // run server
    router.Run()
}
