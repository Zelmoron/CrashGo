package endpoint

import (
	"CaseGo/internal/models"
	"CaseGo/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		//Определяем методы
		GetUsers(int) models.UserModel
		CreateUser(int, string) models.UserModel
		GetInventory(int) []service.Inventory
		GetCases() []service.Cases
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
	response := e.service.CreateUser(user.Id, user.Name)

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"data": response,
	})

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

func (e *Endpoint) GetInventory(c *fiber.Ctx) error {
	var user UserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "BadRequest"})
	}

	response := e.service.GetInventory(user.Id)

	if len(response) < 1 {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": response})
	}

	jsonBytes, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": string(jsonBytes)})

}

func (e *Endpoint) GetCases(c *fiber.Ctx) error {

	response := e.service.GetCases()

	if len(response) < 1 {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": response})
	}

	jsonBytes, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "inv", "data": string(jsonBytes)})

}
