package controllers

import (
	"casego/internal/dto"
	"casego/internal/services"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	//Метод для добавления пользователя в бд или проверки на его нахождение в ней
	var userDTO dto.UserDTO

	if err := c.BodyParser(&userDTO); err != nil { // parsing input data
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	} // получаем из тела запроса информацию о пользователе

	if validationErr := validate.Struct(userDTO); validationErr != nil { // DTO validation
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationErr.Error()})
	} // проверяем валидность

	fmt.Print(userDTO) // просто для теста, можно удалить, но пока не советую

	//TODO
	// Сделать проверку на то, есть ли пользователь в бд

	user, err := services.CreateUser(userDTO, db) // calling business logic from the service
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "user created", "data": user})
}

func GetUsers(c *fiber.Ctx, db *gorm.DB) error {
	users, err := services.GetUsers(db)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "users not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users})
}

func GetUser(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")

	user, err := services.GetUser(id, db)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func UpdateUser(c *fiber.Ctx, db *gorm.DB, validate *validator.Validate) error {
	id := c.Params("id")
	var userDTO dto.UserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if validationErr := validate.Struct(userDTO); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationErr.Error()})
	}

	user, err := services.UpdateUser(id, userDTO, db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "user updated", "data": user})
}

func DeleteUser(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")

	if err := services.DeleteUser(id, db); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "user deleted"})
}
