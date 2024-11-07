package database

import (
	"CaseGo/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func New() *Database {
	return &Database{}
}

func (d *Database) CreateTables() *gorm.DB {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Irkutsk",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error: ", err)

	}

	if err := db.AutoMigrate(&models.UserModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	log.Println("Таблицы успешно созданы")

	return db

}
