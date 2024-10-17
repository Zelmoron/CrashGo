package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	time.Sleep(2 * time.Second)
	return c.Status(200).Render("index", fiber.Map{}) //page of index
}
