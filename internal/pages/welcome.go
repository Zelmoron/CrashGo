package pages

import "github.com/gofiber/fiber/v2"

func WelcomePage(c *fiber.Ctx) error {
	// Rendering page and send fetch js on PostUser, after add or check user can go on next page
	return c.Status(200).Render("welcome", fiber.Map{}) // welvome page
}
