package database

import (
	"log"
	"os"

	"github.com/send2gopal/gocrud/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DBCONNECTION")), &gorm.Config{})

	log.Println("Migrating DB")
	DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}
}

func MigrateDatabase() {
	log.Println("Starting DB Migration")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error while migrating Database", err)
	}
}
