package main

import (
	"casego/internal/db"
	"casego/internal/logic"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var validate *validator.Validate

func main() {

	validate = validator.New() // initializing the validator

	err := godotenv.Load() // loading environment variables
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	database, err := db.Connect() // connecting to the DB

	if err != nil {
		return
	}

	engine := html.New("./views", ".html") // connecting to html

	app := fiber.New(fiber.Config{ // creating a router
		Views: engine,
	})

	app.Use(cors.New())

	logic.Routes(app, database, validate) // starting handlerss

	PORT := os.Getenv("PORT")

	app.Listen(fmt.Sprintf(":%s", PORT)) // listening on port 3000

}
