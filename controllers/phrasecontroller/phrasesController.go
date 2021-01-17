package phrasecontroller

import (
	"math/rand"

	"github.com/danielsoro/amanda/models"
	"github.com/gofiber/fiber/v2"
)

// GetPhrasesHandler :: GET - Handler to get phrases
func GetPhrasesHandler(c *fiber.Ctx) error {
	phrases := models.Phrase{}.GetPhrases()
	bind := fiber.Map{
		"phrase": phrases[rand.Intn(len(phrases))].Value,
	}
	err := c.Render("index", bind)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	return nil
}
