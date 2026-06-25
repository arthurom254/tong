package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v4"
	"github.com/tong/forms"
)

func main() {
	textField := forms.TextField{
		FieldType:  "text",
		FieldName:  "phone",
		FieldValue: "0712345678",
		FieldAttrs: map[string]any{
			"class":       "bg-red-300 mx-auto",
			"placeholder": "enter your phone",
			"required":    true,
		},
	}
	fmt.Println(textField)
	engine := django.New("./views", ".django")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"form": textField,
		})
	})

	app.Listen(":8001")
	// html := tong.Render(textField) {{}}
	// fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
