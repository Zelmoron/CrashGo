package controllers

import (
	"casego/internal/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Inventory struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Type string  `json:"type"`
}

func GetInventroty(c *fiber.Ctx, db *gorm.DB) error {
	//Функкция которая выдает инвентарь
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

	if len(inv) < 1 {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": inv})
	}

	jsonBytes, err := json.Marshal(&inv)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": string(jsonBytes)})

}

func DropItem(c *fiber.Ctx, db *gorm.DB) error {

	//Функция обработчик выпадения придмета,
	//TODO
	//Сделать рандом и отправлять на фронтенд картинку
	id := c.Params("id")

	// var cases CaseRequest

	// if err := c.BodyParser(&cases); err != nil { // parsing input data
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	// } // получаем из тела запроса информацию о id кейса
	var caseData models.Cases
	if err := db.Preload("Items").First(&caseData, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Case not found"})
	}
	fmt.Println(caseData.Items)

	c.JSON(caseData)

	var create = models.InventoryModel{
		TelegramID: uint(1628918728),
		Name:       caseData.Items[1].Name,
		Cost:       float32(caseData.Items[1].Cost),
		Type:       caseData.Items[1].Type,
		Image:      caseData.Items[1].Image,
	}

	if err := db.Create(&create).Error; err != nil {
		return err
	}
	return nil
}
