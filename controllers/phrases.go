package controllers

import (
	"math/rand"

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
		"Não existe ninguém mais linda que você!!!",
	}
}

// GetPhrasesHandler :: GET - Handler to get phrases
func GetPhrasesHandler(ctx *fiber.Ctx) {
	var phrases = getPhrases()
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))],
	}
	ctx.Render("view/index.html", bind)
}
