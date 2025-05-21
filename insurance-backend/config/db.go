package config

import (
	"fmt"
	"insurance-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=johnny password=080906 dbname=insurance port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.Client{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	fmt.Println("DB Connected")
	return db
}
