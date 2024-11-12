package app

import (
	"CaseGo/internal/database"
	"CaseGo/internal/endpoint"
	"CaseGo/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	a.app.Use(cors.New(), logger.New())

	a.service = service.New(db) //получаем с.Service

	a.endpoint = endpoint.New(a.service) //Получаем с.Endpoints( перед этим мы посылаем структуру Service в интерфейс, а также переменную db)

	a.routers() //Вызываем функцию для роутеров

	return a

}

func (a *App) routers() {
	//Здесь хранятся роутеры
	a.app.Get("/users", a.endpoint.GetUsers)             //Метод для получения конкретного пользователя
	a.app.Post("/users", a.endpoint.CreateUser)          //Метод для проверки/добавления  пользователя
	a.app.Get("/inventory/:id", a.endpoint.GetInventory) //Метод для получения инвентаря
	a.app.Get("/cases/cases", a.endpoint.GetCases)       //Метод для получения всех кейсов
	a.app.Get("/cases/weapons/:id", a.endpoint.GetWeapons)
}

func (a *App) Run() {
	//Здесь запускается приложение

	a.app.Listen(":3000")
}
