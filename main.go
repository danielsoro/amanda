package main

import (
	"log"
	"os"

	"github.com/danielsoro/amanda/database"
	"github.com/danielsoro/amanda/routes"

	"github.com/danielsoro/amanda/middlewares"
	"github.com/gofiber/basicauth"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Start the database
	database.Db{}.Session()

	// Create the app
	app := fiber.New()

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
