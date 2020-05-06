package phrasecontroller

import (
	"math/rand"

	"github.com/gofiber/fiber"
)

// GetPhrasesHandler :: GET - Handler to get phrases
func GetPhrasesHandler(ctx *fiber.Ctx) {
	//phrases := models.Phrase{}.GetPhrases()
	phrases := []string{
		"Eu te amo!!",
		"Não vejo a hora de dormir contigo em meus braços todas as noites",
		"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!",
		"Não existe princesa mais linda do que você, nem na Disney",
		"Minha tchutchuquinha, que eu amo muitcho!!!",
		"Já disse que tu é muito linda? Tu é MUITOO LINDA!!",
	}
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))],
	}
	ctx.Render("view/index.html", bind)
}
