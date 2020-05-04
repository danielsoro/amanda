package main

import (
	"github.com/danielsoro/amanda/database"
	"log"
	"os"

	"github.com/danielsoro/amanda/middlewares"
	"github.com/gofiber/basicauth"

	"github.com/danielsoro/amanda/controllers"
	"github.com/gofiber/fiber"
)

func main() {
	// Start the database
	database.Connect()

	// Create the app
	app := fiber.New(&fiber.Settings{
		Prefork: true,
	})

	// Set the basich auth
	app.Use(basicauth.New(middlewares.GetConfig()))

	// Creat the root path
	app.Get("/", controllers.GetPhrasesHandler)

	// Get the port from env
	port := os.Getenv("PORT")

	// verify if port is setted or not
	if port == "" {
		port = "8080"
		log.Print("$PORT == 8080")
	}

	// start the server
	app.Listen(port)
}
