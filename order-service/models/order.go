package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId uint
	FoodId uint
	Status string // "pending", "confirmed", etc.
}
