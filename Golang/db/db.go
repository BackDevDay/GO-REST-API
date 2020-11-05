package db

import (
	"go-rest-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func getDatabase() *gorm.DB {
	return database
}

// ConnectDB - connect DB
func ConnectDB() {
	// set mySQL
	dsn := "root:12345678@tcp(0.0.0.0:9001)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	database = db
}
