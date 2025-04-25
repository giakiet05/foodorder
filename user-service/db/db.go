package db

import (
	"github.com/giakiet05/foodorder/user-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB.AutoMigrate(&models.User{})
	log.Println("Connected to DB and migrated")
}
