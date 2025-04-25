package models

import "gorm.io/gorm"

// Cấu trúc user sẽ tự ánh xạ thành bảng trong DB
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
