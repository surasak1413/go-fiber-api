package routes

import (
	"loly/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", handler.RegisterHandler)
}
