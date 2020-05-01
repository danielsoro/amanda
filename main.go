package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

func getPhrases() []string {
	return []string{
		"Eu te amo!!",
		"Não vejo a hora de dormir contigo em meus braços todas as noites",
		"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!",
		"Não existe princesa mais linda do que você, nem na Disney",
		"Minha tchutchuquinha, que eu amo muitcho!!!",
		"Já disse que tu é muito linda? Tu é MUITOO LINDA!!",
	}
}

func handler(ctx *fiber.Ctx) {
	var phrases = getPhrases()
	ctx.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	ctx.Send(fmt.Sprintf("<h1>%s</h1>", phrases[rand.Intn(len(phrases))]))
}

func main() {
	// Create the app
	app := fiber.New(&fiber.Settings{
		Prefork: true,
	})

	// Configure basich auth for the project
	cfg := basicauth.Config{
		Users: map[string]string{
			os.Getenv("USER"): os.Getenv("PASSWORD"),
		},
	}

	// Set the basich auth
	app.Use(basicauth.New(cfg))

	// Creat the root path
	app.Get("/", handler)

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
