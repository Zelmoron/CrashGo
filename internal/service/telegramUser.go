package service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Coins int
}

func GetTelegramUser(c *fiber.Ctx) error {
	var u User

	fmt.Println("Test")

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	} // parsde data from fetch

	return c.JSON(fiber.Map{"status": "success", "data": u}) // response
}
