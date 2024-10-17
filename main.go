package main

import (
	"bot/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html") // connect html

	app := fiber.New(fiber.Config{
		Views: engine,
	}) // create router

	service.Hadlers(app) // start handlers

	app.Listen(":3000") // listen on 3000 port

}
