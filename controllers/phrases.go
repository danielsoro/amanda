package controllers

import (
	"math/rand"

	"github.com/danielsoro/amanda/database"
	"github.com/danielsoro/amanda/models"

	"github.com/gofiber/fiber"
)

func getPhrases() []models.Phrase {
	var phrases []models.Phrase
	database.Connect().Find(&phrases)
	return phrases
}

// GetPhrasesHandler :: GET - Handler to get phrases
func GetPhrasesHandler(ctx *fiber.Ctx) {
	var phrases = getPhrases()
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))].Value,
	}
	ctx.Render("view/index.html", bind)
}
