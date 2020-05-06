package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielsoro/amanda/database"
	"github.com/danielsoro/amanda/routes"

	"github.com/danielsoro/amanda/middlewares"
	"github.com/gofiber/basicauth"
	"github.com/gofiber/recover"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Start the database
	database.Session()

	// Create the app
	app := fiber.New()

	// Configure Middlewares
	app.Use(basicauth.New(middlewares.GetConfig()))
	app.Use(recover.New())

	// Creating routes
	routes.Routes(app)

	// 404 Middleware
	app.Use(middlewares.NotFoundHandler)

	// Get the port from env
	port := os.Getenv("PORT")

	// verify if port is setted or not
	if port == "" {
		port = "8080"
		log.Print("$PORT == 8080")
	}

	// start the server
	if err := app.Listen(port); err != nil {
		fmt.Errorf("error during server startup %v", err.Error())
	}

}
