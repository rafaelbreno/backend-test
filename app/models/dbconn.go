package models

import (
	"coderockr-test/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	mig := database.Conn()

	mig.AutoMigrate(&User{})

	DB = mig
}
