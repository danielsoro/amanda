package controllers

import (
	"math/rand"

	"github.com/danielsoro/amanda/models"
	"github.com/gofiber/fiber"
)

// PhraseController is the struct for phrase
type PhraseController struct {
}

// GetPhrasesHandler :: GET - Handler to get phrases
func (p PhraseController) GetPhrasesHandler(ctx *fiber.Ctx) {
	phrases := models.Phrase{}.GetPhrases()
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))].Value,
	}
	ctx.Render("view/index.html", bind)
}
