package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	Service interface {
		//Определяем методы
		GetUsers(database *gorm.DB)
	}

	Endpoint struct {
		service Service  //поле для интерфейса
		db      *gorm.DB //поле для базы данных
	}
)

func New(service Service, db *gorm.DB) *Endpoint {
	//Констркутор для с. Endpoint
	return &Endpoint{
		service: service, //Передали структуру Service, и поместили в поле service типа Service(Интерфейс)
		db:      db,      // Поле для бд
	}
}

func (e *Endpoint) GetUsers(c *fiber.Ctx) error {
	//Метод для получение пользователя
	e.service.GetUsers(e.db)
	return nil
}
