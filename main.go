package main

import (
	"log"
	"os"

	"github.com/danielsoro/amanda/database"
	"github.com/danielsoro/amanda/routes"

	"github.com/danielsoro/amanda/middlewares"
	"github.com/gofiber/basicauth"

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

	// Creating routes
	routes.Routes(app)

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
