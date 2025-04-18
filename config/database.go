package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"finalproject/models"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Simpan ke global DB
	DB = db

	// Auto migrate semua model
	err = db.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderDetail{},
	)
	if err != nil {
		log.Fatal("Failed to auto migrate models:", err)
	}

	fmt.Println("Database connected and migrated!")
}
