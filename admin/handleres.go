package admin

import "github.com/gofiber/fiber/v2"

func DashboarboardHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
