package db

import (
	"github.com/giakiet05/foodorder/food-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("food.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to DB", err)
	}

	DB.AutoMigrate(&models.Food{})
	log.Println("Connected to food DB")
}
