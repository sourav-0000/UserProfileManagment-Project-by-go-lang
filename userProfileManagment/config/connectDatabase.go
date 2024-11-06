package config

import (
	"log"
	"userProfileManagment/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=root dbname=employee port=9920 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	db.AutoMigrate(model.User{})
	DB = db
	return db, nil
}
