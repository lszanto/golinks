package main

import (
	"encoding/json"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lszanto/links/config"
	"github.com/lszanto/links/models"
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

	// migrate schema
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Link{})
}
