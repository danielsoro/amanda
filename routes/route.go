package routes

import (
	"github.com/danielsoro/amanda/controllers/phrasecontroller"
	"github.com/gofiber/fiber/v2"
)

// Routes is a function to define all routes of the project
func Routes(app *fiber.App) {
	app.Get("/", phrasecontroller.GetPhrasesHandler)
}
