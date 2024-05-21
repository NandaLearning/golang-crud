package config

import (
	"latihan-go-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=belajargo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Siswa{})
}

func GetDB() *gorm.DB {
	return db
}
