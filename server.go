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

	// SETUP CONTROLLERS

	// controllers
	lc := controllers.NewLinkController(db, config)
	uc := controllers.NewUserController(db, config)

	// ROUTER

	// setup router
	router := gin.Default()

	// login routes
	router.POST("/user/login", uc.Login)
	router.POST("/user", uc.CreateUser)
	router.GET("/user/:id", uc.Get)

	// link routes
	router.GET("/links", lc.GetAll)
	router.GET("/links/:id", lc.Get)
	router.PUT("/links/:id", middleware.JWTVerify(config.SecretKey), lc.Update)
	router.DELETE("/links/:id", middleware.JWTVerify(config.SecretKey), lc.Delete)
	router.POST("/links", middleware.JWTVerify(config.SecretKey), lc.Create)

	// SET STATIC DIR, START SERVER

	// setup static folder
	router.Static("/assets", "./assets")

	// run server
	router.Run()
}
