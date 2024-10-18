package service

import (
	"bot/internal/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {

	//Get request from frontend and start work with datavase

	var u db.User
	log.Println("Создание или проверка на наличие в базе данных")

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	} // parsde data from fetch

	err := u.CreateUser()

	if err != nil {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{"status": "success", "data": u}) // response
}
