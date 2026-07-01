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
		var Form_ forms.Form = forms.NewModelFrom(&User{}, "form-input px-2")
		return c.Render("index", fiber.Map{
			"form":  RegForm,
			"Form_": Form_,
		})
	})
	app.Post("/", func(c *fiber.Ctx) error {
		var Form_ forms.Form = forms.NewModelFrom(&User{}, "form-input px-2")

		form := Form_
		var user User
		if form.Bind(c, &user) {
			fmt.Println(user.Email)
			fmt.Println(user.FirstName, user.LastName)
			return c.SendString("OK ")
		}

		return c.SendString("OK")
	})

	app.Listen(":8001")
	// html := tong.Render(textField) {{}}
	// fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
