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

	if err := db.AutoMigrate(&models.InventoryModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	if err := db.AutoMigrate(&models.ItemModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	if err := db.AutoMigrate(&models.CasesModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	// Сначала создаем все кейсы
	var case1, case2, case3 models.CasesModel

	// Проверяем, существует ли уже кейс "Решающий момент"
	result := db.Where("name = ?", "Решающий момент").First(&case1)
	if result.Error != nil {
		// Если нет, создаем новый
		case1 = models.CasesModel{Name: "Решающий момент", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/decisive_moment.png"}
		db.Create(&case1)
	}

	// Проверяем, существует ли уже кейс "Гидра"
	result = db.Where("name = ?", "Гидра").First(&case2)
	if result.Error != nil {
		// Если нет, создаем новый
		case2 = models.CasesModel{Name: "Гидра", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/hydra.png"}
		db.Create(&case2)
	}

	// Проверяем, существует ли уже кейс "Фальшион"
	result = db.Where("name = ?", "Фальшион").First(&case3)
	if result.Error != nil {
		// Если нет, создаем новый
		case3 = models.CasesModel{Name: "Фальшион", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/falchion.png"}
		db.Create(&case3)
	}

	// Теперь создаем элементы, связанные с кейсами
	items := []models.ItemModel{
		{Name: "Item 1.1", Cost: 1000, Type: "gun", Image: "/", CaseID: case1.ID},
		{Name: "Item 1.2", Cost: 2000, Type: "gun", Image: "/", CaseID: case1.ID},
		{Name: "Item 2.1", Cost: 3000, Type: "gun", Image: "/", CaseID: case2.ID},
		{Name: "Item 3.1", Cost: 3000, Type: "gun", Image: "/", CaseID: case3.ID},
		{Name: "Item 3.2", Cost: 4000, Type: "gun", Image: "/", CaseID: case3.ID},
	}

	for _, item := range items {
		db.Create(&item)
	}

	log.Println("Таблицы успешно созданы")

	return db

}
