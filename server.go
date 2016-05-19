package main

import (
    "encoding/json"
    "os"

    "github.com/gin-gonic/gin"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/lszanto/links/config"
    "github.com/lszanto/links/controllers"
    "github.com/lszanto/links/middleware"
)

func main() {
    // set error holder
    var err error

    // CONFIG

    // create config object
    config := config.Config{}

    // open a config file
    configFile, _ := os.Open("app/config.json")

    // decode into config
    err = json.NewDecoder(configFile).Decode(&config)

    if err != nil {
        panic("failed to open config")
    }

    // GORM DATABASE

    // create db connection
    db, err := gorm.Open(config.DatabaseEngine, config.DatabaseString)

    if err != nil {
        panic("failed to connect to database")
    }

    // controllers
    lc := controllers.NewLinkController(db, config)
    uc := controllers.NewUserController(db, config)

    // ROUTER

    // setup router
    router := gin.Default()

    // login routes
    router.POST("/login", uc.Login)

    // link routes
    router.GET("/link/:id", lc.Get)
    router.DELETE("/link/:id", middleware.JWTVerify(config.SecretKey), lc.Delete)
    router.POST("/link", middleware.JWTVerify(config.SecretKey), lc.Post)

    // SET STATIC DIR, START SERVER

    // setup static folder
    router.Static("/assets", "./assets")

    // run server
    router.Run()
}
