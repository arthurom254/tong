package admin

import "github.com/gofiber/fiber/v2"

func SetupRoutes(admin fiber.Router) error {
	admin.Get("/dashboard", DashboarboardHandler)
	return nil
}
