package logic

import (
	"casego/internal/controllers"
	"casego/internal/pages"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB, validate *validator.Validate) {

	app.Static("/static", "./static") // connection css and js

	app.Get("/", pages.WelcomePage) // rendering the welcome page
	app.Get("/api", pages.Index)    // rendering the main page

	// CRUD User
	app.Post("/users", func(c *fiber.Ctx) error {
		log.Println(c.IP())

		return controllers.CreateUser(c, db, validate)
	})
	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetUsers(c, db)
	})
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return controllers.GetUser(c, db)
	})
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateUser(c, db, validate)
	})
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteUser(c, db)
	})

	// app.Post("/random", func(c *fiber.Ctx) error {
	// 	return controllers.GetRandomNumber(c, db)
	// })

	app.Get("/inventory/:id", func(c *fiber.Ctx) error {
		return controllers.GetInventroty(c, db)
	})

	app.Get("/dropitem/:id", func(c *fiber.Ctx) error {
		return controllers.DropItem(c, db)
	})

	// app.ListenTLS(":3000", "localhost.crt", "localhost.key")
}
