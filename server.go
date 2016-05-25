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
	sc := controllers.NewSiteController(db, config)

	// ROUTER

	// setup router
	router := gin.Default()

	// let router use CORSMiddleware
	router.Use(middleware.CORSMiddleware())

	// load templates
	router.LoadHTMLGlob("templates/*")

	// base routes
	router.GET("/", sc.Home)

	// api routes
	api := router.Group("/api")
	{
		// login routes anybody
		api.POST("/user/login", uc.Login)
		api.POST("/user", uc.CreateUser)
		api.GET("/user/:id", uc.Get)

		// link routes anybody
		api.GET("/links", lc.GetAll)
		api.GET("/links/:id", lc.Get)

		// set auth routes
		auth := api.Group("/", middleware.JWTVerify(config.SecretKey))
		{
			auth.PUT("/links/:id", lc.Update)
			auth.DELETE("/links/:id", lc.Delete)
			auth.POST("/links", lc.Create)
		}
	}

	// SET STATIC DIR, START SERVER

	// setup static folder
	router.Static("/assets", "./assets")

	// run server
	router.Run(":3000")
}
