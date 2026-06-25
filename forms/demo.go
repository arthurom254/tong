package forms

var LoginForm = Form{
	Fields: []Widget{
		TextField{
			FieldType:  "email",
			FieldName:  "email",
			FieldValue: "",
			FieldAttrs: map[string]any{
				"class":       "bg-red-300 mx-auto",
				"placeholder": "enter your email",
				"required":    true,
			},
		},
		TextField{
			FieldType:  "password",
			FieldName:  "password",
			FieldValue: "",
			FieldAttrs: map[string]any{
				"class":       "bg-red-300 mx-auto",
				"placeholder": "enter your password",
				"required":    true,
			},
		},
	},
}
