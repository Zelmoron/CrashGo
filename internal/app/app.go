package app

import (
	"CaseGo/internal/database"
	"CaseGo/internal/endpoint"
	"CaseGo/internal/service"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	//Стукрутра приложения
	app      *fiber.App         //Поле для fiber
	endpoint *endpoint.Endpoint //Поле для стуктуры Endpoints
	service  *service.Service   //Поле для с. Service
	database *database.Database //Поле для с. Database
}

func New() *App {
	a := &App{} // Создаем экземпляр

	a.database = database.New() //Получаем с. Database

	db := a.database.CreateTables() //Получаем переменую тип *gorm.DB из метода структуры Database

	a.app = fiber.New() //создаем приложение на Fiber

	a.service = service.New() //получаем с.Service

	a.endpoint = endpoint.New(a.service, db) //Получаем с.Endpoints( перед этим мы посылаем структуру Service в интерфейс, а также переменную db)

	a.routers() //Вызываем функцию для роутеров

	return a

}

func (a *App) routers() {
	//Здесь хранятся роутеры
	a.app.Get("/users", a.endpoint.GetUsers) //Метод для получения конкретного пользователя
}

func (a *App) Run() {
	//Здесь запускается приложение

	a.app.Listen(":3000")
}
