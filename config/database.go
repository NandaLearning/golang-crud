package config

import (
	"latihan-go-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "user=postgres.kznzozphkyddljcrgvjj password=cAl9QgmgxxID77m1 host=aws-0-ap-southeast-1.pooler.supabase.com port=5432 dbname=postgres"
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
