package middlewares

import "github.com/gofiber/fiber"

// NotFoundHandler Render the default NotFound page
func NotFoundHandler(c *fiber.Ctx) {
	c.Status(404)
	err := c.Render("template/view/notfound.html", nil)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}
}
