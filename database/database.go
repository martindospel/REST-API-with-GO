package database

import (
	"log"
	"os"

	"github.com/martindospel/REST-API-with-GO.git/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//this will create tables for user, products and orders
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
