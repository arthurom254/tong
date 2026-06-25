package main

import (
	"fmt"

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
	// html := tong.Render(textField) {{}}
	// fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
