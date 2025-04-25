package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}
