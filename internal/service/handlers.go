package service

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Coins int
}

func Hadlers(app *fiber.App) {

	app.Static("/static", "./static") // connection css and js

	app.Get("/", Index)
	app.Post("/api/data", GetTelegramUser)

	// app.ListenTLS(":3000", "localhost.crt", "localhost.key")

}
