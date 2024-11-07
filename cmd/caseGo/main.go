package main

import (
	"CaseGo/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	//В инициализаторе качаем данные из env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}
func main() {

	app := app.New() // получаем структуру App
	app.Run()        // Запускаем приложение
}
