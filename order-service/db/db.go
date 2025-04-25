package db

import (
	"github.com/giakiet05/foodorder/order-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect DB", err)
	}
	DB.AutoMigrate(&models.Order{})
	log.Println("Connected to Order DB")

}
