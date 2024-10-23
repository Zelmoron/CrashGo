package db

import (
	"casego/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Irkutsk",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error: ", err)
		return nil, err
	}

	if err := db.AutoMigrate(&models.UserModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	if err := db.AutoMigrate(&models.ItemsModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	if err := db.AutoMigrate(&models.InventoryModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	return db, nil
}
