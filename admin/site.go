package admin

import "github.com/gofiber/fiber/v2"

func SetupSite(app *fiber.App, site_name string, site_path string) error {
	if site_path == "" {
		site_name = "/admin"
	}
	if site_name == "" {
		site_name = "Tong Admin"
	}
	admin := app.Group(site_name)

	SetupRoutes(admin)
	return nil
}
