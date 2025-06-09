package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"insurance-backend/models"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=johnny password=080906 dbname=insurance port=5432 sslmode=disable"
	fmt.Println("Trying to connect to:", dsn)
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
