package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Initialize the database connection
	dns := "root:123456@tcp(127.0.0.1:3306)/db_blogs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = db
	fmt.Println("Database connected!")
}
