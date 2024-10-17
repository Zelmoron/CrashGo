package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Coins int
}

var data = make(map[int]int)

func main() {

	data[1628918728] = 124 // типо база данных, пока что

	engine := html.New("./views", ".html") // connect html

	app := fiber.New(fiber.Config{
		Views: engine,
	}) // create router

	app.Static("/static", "./static") // connection css and js

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(2 * time.Second)
		return c.Status(200).Render("index", fiber.Map{}) //page of index
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).Render("download", fiber.Map{}) //page of index
	})

	app.Post("/api/data", func(c *fiber.Ctx) error {

		var u User

		fmt.Println("Test")

		if err := c.BodyParser(&u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		u.Coins = data[u.Id]
		// Обработка данных
		return c.JSON(fiber.Map{"status": "success", "data": u})
	})

	// app.ListenTLS(":3000", "localhost.crt", "localhost.key")

	app.Listen(":3000")

}
