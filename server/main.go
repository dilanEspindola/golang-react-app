package main

import (
	"anime-hentai-backend/config"
	"anime-hentai-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	api := app.Group("api/")
	api.Get("users", routes.GetUsers)
	api.Post("users", routes.CreateUser)
}

func main() {
	app := fiber.New()
	config.DbConnection()

	setUpRoutes(app)

	app.Listen(":4000")
}
