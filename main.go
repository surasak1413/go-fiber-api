package main

import (
	"loly/database"
	"loly/repository"
	"loly/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		return
	}
	repository.NewStore(db)
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":7777")
}
