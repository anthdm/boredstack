package handlers

import "github.com/gofiber/fiber/v2"

func HandleHome(c *fiber.Ctx) error {
	return c.Render("home/index", fiber.Map{})
}
