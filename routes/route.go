package routes

import (
	"github.com/danielsoro/amanda/controllers"
	"github.com/gofiber/fiber"
)

// Routes is a function to define all routes of the project
func Routes(app *fiber.App) {
	app.Get("/", controllers.GetPhrasesHandler)
}
