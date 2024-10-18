package service

import (
	"github.com/gofiber/fiber/v2"
)

func Hadlers(app *fiber.App) {

	app.Static("/static", "./static") // connection css and js

	app.Get("/", WelcomePage) // rendering welcome page

	app.Get("/app", Index) // rendering main page

	usersHadler := app.Group("") //CRUD for users

	usersHadler.Post("/api/data", PostUser) // create Users if not exists

	// usersHadler.Get("/api/data", GetUser) // get Users if not exists

	// usersHadler.Patch("/api/data", UpdateUser) // update Users if not exists

	// usersHadler.Delete("/api/data", DeleteUser) // delete Users if not exists

	// app.ListenTLS(":3000", "localhost.crt", "localhost.key")

}
