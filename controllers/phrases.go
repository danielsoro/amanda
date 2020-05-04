package controllers

import (
	"math/rand"

	"github.com/danielsoro/amanda/models"
	"github.com/gofiber/fiber"
)

// GetPhrasesHandler :: GET - Handler to get phrases
func GetPhrasesHandler(ctx *fiber.Ctx) {
	var phrases = models.GetPhrases()
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))].Value,
	}
	ctx.Render("view/index.html", bind)
}
