package db

import (
	"rest/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

// GetDatabase - for call
func GetDatabase() *gorm.DB {
	return database
}

// ConnectDB - connect DB
func ConnectDB() {
	// set mySQL
	dsn := "root:12345678@tcp(mysql:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})

	database = db
}
