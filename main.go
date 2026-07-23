package main

import (
	"fmt"

	"github.com/arthurom254/tong/admin"
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
			"form":      RegForm,
			"Form_":     Form_,
			"LoginForm": forms.NewModelFrom(&Login{}, ""),
		})
	})
	app.Post("/", func(c *fiber.Ctx) error {
		// var Form_ forms.Form = forms.NewModelFrom(&User{}, "form-input px-2")
		var LoginForm forms.Form = forms.NewModelFrom(&Login{}, "form-input px-2")

		form := LoginForm
		var user Login
		if form.Bind(c, &user) {
			fmt.Println(user.Username)
			fmt.Println(user.Password)
			return c.SendString("OK3")
		}

		return c.SendString(fmt.Sprint(form.Errors))
	})

	// app.Route("/admin", )
	admin.SetupSite(app, "", "")

	app.Listen(":8001")
	// html := tong.Render(textField) {{}}
	// fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
