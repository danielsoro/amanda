package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber"
)

type Phrases struct {
	value string
}

func getPhrases() []Phrases {
	return []Phrases{
		{"Eu te amo!!"},
		{"Não vejo a hora de dormir contigo em meus braços todas as noites"},
		{"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!"},
		{"Não existe princesa mais linda do que você, nem na Disney"},
		{"Minha tchutchuquinha, que eu amo muitcho!!!"},
		{"Já disse que tu é muito linda? Tu é MUITOO LINDA!!"},
	}
}

func handler(ctx *fiber.Ctx) {
	var phrases = getPhrases()
	ctx.Set("Content-Type", "text/html")
	ctx.Send(fmt.Sprintf("<h1>%s</h1>", phrases[rand.Intn(len(phrases))].value))
}

func main() {
	app := fiber.New()

	app.Get("/", handler)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Print("$PORT == 8080")
	}

	app.Listen(port)
}
