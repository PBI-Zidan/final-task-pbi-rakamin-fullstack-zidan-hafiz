package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=hafiz211001 dbname=go-rest port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{}, &Photo{})

	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = db
}
