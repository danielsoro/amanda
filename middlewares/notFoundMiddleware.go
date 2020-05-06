package middlewares

import "github.com/gofiber/fiber"

func NotFoundHandler(c *fiber.Ctx) {
	c.Status(404)
	err := c.Render("view/notfound.html", nil)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}
}
