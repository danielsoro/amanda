package main

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
	"log"
	"os"

	"github.com/danielsoro/amanda/database"
	"github.com/danielsoro/amanda/middlewares"
	"github.com/danielsoro/amanda/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	xForwardedProtoHeader = "x-forwarded-proto"
)

func main() {
	// Start the database
	database.Session()


	// Create the template engine
	engine := html.New("./template/view", ".html")

	// Create the app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static files
	app.Static("/", "./public")

	// Configure Middlewares
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
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
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("error during server startup %v", err.Error())
	}
}
