package service

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	return c.Status(200).Render("index", fiber.Map{}) //page of index
}
