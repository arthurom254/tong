package forms

var passwordField = PasswordField{
	Value:     "",
	BaseField: BaseField{},
}
var LoginForm = Form{
	Fields: []Widget{
		TextField{
			FieldType:  "label",
			FieldName:  "month",
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
