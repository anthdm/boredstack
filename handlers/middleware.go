package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).Render("error/404", nil)
}

func FlashMiddleware(c *fiber.Ctx) error {
	c.Locals("flash", flash.Get(c))
	return c.Next()
}
