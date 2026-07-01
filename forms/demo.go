package forms

import (
	"time"
)

var inputClass = ""

var RegForm = Form{
	Fields: []Field{
		CharField{
			Name:     "first_name",
			Label:    "First Name",
			Required: true,
			HelpText: "Enter your given name.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "first_name",
				"name":         "first_name",
				"type":         "text",
				"class":        inputClass,
				"placeholder":  "John",
				"autocomplete": "given-name",
				"maxlength":    100,
			},
		},
		CharField{
			Name:     "last_name",
			Label:    "Last Name",
			Required: true,
			HelpText: "Enter your family name.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "last_name",
				"name":         "last_name",
				"type":         "text",
				"class":        inputClass,
				"placeholder":  "Doe",
				"autocomplete": "family-name",
				"maxlength":    100,
			},
		},
		CharField{
			Name:     "username",
			Label:    "Username",
			Required: true,
			HelpText: "Choose a unique username.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "username",
				"name":         "username",
				"type":         "text",
				"class":        inputClass,
				"placeholder":  "johndoe",
				"autocomplete": "username",
				"maxlength":    150,
			},
		},
		CharField{
			Name:     "email",
			Label:    "Email Address",
			Required: true,
			HelpText: "We'll never share your email.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "email",
				"name":         "email",
				"type":         "email",
				"class":        inputClass,
				"placeholder":  "john@example.com",
				"autocomplete": "email",
			},
		},
		CharField{
			Name:     "phone",
			Label:    "Phone Number",
			HelpText: "Optional.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "phone",
				"name":         "phone",
				"type":         "tel",
				"class":        inputClass,
				"placeholder":  "+254712345678",
				"autocomplete": "tel",
			},
		},
		CharField{
			Name:     "password",
			Label:    "Password",
			Required: true,
			HelpText: "Must be at least 8 characters.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "password",
				"name":         "password",
				"type":         "password",
				"class":        inputClass,
				"placeholder":  "Password",
				"autocomplete": "new-password",
				"minlength":    8,
			},
		},
		CharField{
			Name:     "confirm_password",
			Label:    "Confirm Password",
			Required: true,
			HelpText: "Re-enter your password.",
			Errors:   []string{},
			Attrs: map[string]any{
				"id":           "confirm_password",
				"name":         "confirm_password",
				"type":         "password",
				"class":        inputClass,
				"placeholder":  "Confirm Password",
				"autocomplete": "new-password",
				"minlength":    8,
			},
		},
	},
}

type User struct {
	FirstName string `json:"first_name" form:"label=First Name;placeholder=John;autocomplete=given-name;maxlength=100;required"`
	LastName  string `json:"last_name" form:"label=Last Name;placeholder=Doe;autocomplete=family-name;maxlength=100;required"`
	Username  string `json:"username" form:"placeholder=johndoe;autocomplete=username;maxlength=150;required"`
	Email     string `json:"email" form:"widget=email;placeholder=john@example.com;autocomplete=email;required"`
	Phone     string `json:"phone" form:"widget=tel;placeholder=+254712345678;autocomplete=tel"`
	Password  string `json:"-" form:"type=password;placeholder=Password;minlength=8;required;exclude_prefill"`

	Bio      string    `json:"bio" form:"widget=textarea;label=About You;help=Tell us a bit about yourself;maxlength=500"`
	Website  string    `json:"website" form:"widget=url;placeholder=https://example.com"`
	Birthday time.Time `json:"birthday" form:"widget=date;label=Date of Birth"`

	IsActive bool   `json:"is_active" form:"widget=checkbox;label=Active Account"`
	Role     string `json:"role" form:"widget=select;label=Role;required"`

	AvatarURL string `json:"avatar_url" form:"-"`

	OrganizationID uint `json:"organization_id" form:"-"`
}

var Form_ Form = NewModelFrom(&User{}, "form-input px-2")
