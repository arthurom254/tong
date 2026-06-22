package main

import (
	"fmt"

	tongforms "github.com/tong/tong_forms"
)

func main() {
	textField := tongforms.TextField{
		FieldType:  "text",
		FieldName:  "phone",
		FieldValue: "0712345678",
		FieldAttrs: map[string]any{
			"class":       "bg-red-300 mx-auto",
			"placeholder": "enter your phone",
		},
	}

	html := tongforms.Render(textField)
	fmt.Println(html) // out: <input type="text" name="phone" placeholder ="enter your phone" class ="bg-red-300 mx-auto">
}
