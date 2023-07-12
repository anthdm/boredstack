package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

func HandleHome(c *fiber.Ctx) error {
	return c.Render("home/index", fiber.Map{})
}

func HandleBored(c *fiber.Ctx) error {
	return c.Render("home/bored", fiber.Map{})
}

func HandleFlash(c *fiber.Ctx) error {
	context := fiber.Map{
		"systemMessage": "a flash message for you user",
	}
	return flash.WithData(c, context).RedirectBack("/")
}
