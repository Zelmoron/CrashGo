package controllers

import (
	"casego/internal/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ItemRequest struct {
	ID         int `json:"id"`
	TelegramID int `json:"telegramid"`
}

type InventoryRequest struct {
	TelegramID int `json:"telegramid"`
}

type Inventory struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Type string  `json:"type"`
}

func GetRandomNumber(c *fiber.Ctx, db *gorm.DB) error {

	var itemResp ItemRequest                        //сюда подается id телеграмма и номер выпавший
	if err := c.BodyParser(&itemResp); err != nil { // parsing input data
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	} // получаем из тела запроса информацию о пользователе

	fmt.Println(itemResp) // просто для теста, можно удалить, но пока не совету

	// позже раскидаю по файлам конекты

	var item models.ItemsModel

	if err := db.Where("id=?", itemResp.ID).First(&item).Error; err != nil {
		return err
	}

	var create models.InventoryModel

	create.TelegramID = uint(itemResp.TelegramID)
	create.Name = item.Name
	create.Cost = item.Cost
	create.Type = item.Type

	if err := db.Create(&create).Error; err != nil {
		return err
	}
	return nil
}

func GetInventroty(c *fiber.Ctx, db *gorm.DB) error {

	id := c.Params("id")
	fmt.Println(id)

	var inventory []models.InventoryModel

	if err := db.Where("telegram_id=?", id).Find(&inventory).Error; err != nil {
		return err
	}

	var inv []Inventory

	for _, v := range inventory {
		inv = append(inv, Inventory{v.Name, v.Cost, v.Type})
	}

	jsonBytes, err := json.Marshal(&inv)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": string(jsonBytes)})

}
