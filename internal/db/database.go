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

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	if err := db.AutoMigrate(&models.InventoryModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	if err := db.AutoMigrate(&models.Cases{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	// case1 := models.Cases{Name: "Case 1"}
	// case2 := models.Cases{Name: "Case 2"}
	// case3 := models.Cases{Name: "Case 3"}

	// db.Create(&case1)
	// db.Create(&case2)
	// db.Create(&case3)

	// items := []models.Item{
	// 	{Name: "Item 1.1", Cost: 1000, Type: "gun", Image: "/", CaseID: case1.ID},
	// 	{Name: "Item 1.2", Cost: 2000, Type: "gun", Image: "/", CaseID: case1.ID},
	// 	{Name: "Item 2.1", Cost: 3000, Type: "gun", Image: "/", CaseID: case2.ID},
	// 	{Name: "Item 3.1", Cost: 3000, Type: "gun", Image: "/", CaseID: case3.ID},
	// 	{Name: "Item 3.2", Cost: 4000, Type: "gun", Image: "/", CaseID: case3.ID},
	// }

	// for _, item := range items {
	// 	db.Create(&item)
	// }

	return db, nil
}
