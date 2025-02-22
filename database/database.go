package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the connection to the MySQL database and returns the DB instance
func ConnectDB() *gorm.DB {
	// Ubah sesuai konfigurasi MySQL Anda
	dsn := "root:@tcp(127.0.0.1:3306)/PusatData?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Open the database connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully!")
	return DB // Return the *gorm.DB instance
}
