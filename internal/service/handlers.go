package service

import (
	"github.com/gofiber/fiber/v2"
)

func Hadlers(app *fiber.App) {

	app.Static("/static", "./static") // connection css and js

	app.Get("/", Index)                    // rendering main page
	app.Post("/api/data", GetTelegramUser) // get information about users

	// app.ListenTLS(":3000", "localhost.crt", "localhost.key")

}
