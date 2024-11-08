package endpoint

import (
	"CaseGo/internal/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		//Определяем методы
		GetUsers(int) models.UserModel
		CreateUser()
	}

	Endpoint struct {
		service Service //поле для интерфейса

	}
)

type UserRequest struct {
	Name string `json:"name" ` // имя пользователя
	Id   int    `json:"id"`    // айди телеграмма!!!!

}

func New(service Service) *Endpoint {
	//Возвращаем с . Endpoint
	return &Endpoint{
		service: service, //Передали структуру Service, и поместили в поле service типа Service(Интерфейс)

	}
}

func (e *Endpoint) CreateUser(c *fiber.Ctx) error {
	var user UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "BadRequest"})
	}
	fmt.Println(user)
	return nil

}
func (e *Endpoint) GetUsers(c *fiber.Ctx) error {
	//Метод для получение пользователя
	var user UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "BadRequest"})
	}
	response := e.service.GetUsers(user.Id)

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"data": response,
	})
}
