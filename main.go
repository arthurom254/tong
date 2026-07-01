package main

import (
	"fmt"

	"github.com/arthurom254/tong/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v4"
)

func main() {
	engine := django.New("./views", ".django")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"form": forms.RegForm,
		})
	})
	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))

		type Result struct {
			Email string `json:"email"`
		}
		var p any
		if err := c.BodyParser(p); err != nil {
			fmt.Println(err)
		}

		fmt.Println("P:-----", p)
		return c.SendString("OK")
	})

	app.Listen(":8001")
	// html := tong.Render(textField) {{}}
	// fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
