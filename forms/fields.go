package forms

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type Field interface {
	getContext() map[string]any
	renderWidget() string
	render() string
	labelTag() string
	getID() any
	GetValue() any
}

type BaseField struct {
}

// text

type CharField struct {
	BaseField
	Name        string
	Type        string
	Errors      []string
	Required    bool
	UseFieldset bool
	Label       string
	HelpText    string
	Disabled    bool
	Value       string
	EmptyValue  string
	Attrs       map[string]any `json:"attr"`
}

func (c CharField) AriaDescribedBy() string {
	if value, ok := c.Attrs["aria_describedby"]; ok {
		return fmt.Sprint(value)
	} else {
		return ""
	}
}

func (c CharField) AsDiv() string {
	return ""
}
func (c CharField) String() string {
	return c.LabelTag() + c.renderWidget()
}

func (c CharField) renderWidget() string {
	ctx := c.getContext()
	name := "widgets/input.html"
	tpl, err := renderTPL(name, ctx)
	if err != nil {
		panic(err)
	}
	return tpl
}

func (c CharField) ToGo() string {
	if c.Value != "" {
		value := string(c.Value)
		return value
	}
	return c.EmptyValue
}

func (c CharField) getContext() pongo2.Context {

	ctx := pongo2.Context{
		"widget": map[string]any{
			"type":             c.Type,
			"name":             c.Name,
			"value":            c.Value,
			"attrs":            c.Attrs,
			"use_fieldset":     c.UseFieldset,
			"aria_describedby": c.AriaDescribedBy, // Call?
			"errors":           c.Errors,
			"label_tag":        c.LabelTag, // Call ?
		},
	}

	return ctx
}

func (c CharField) getID() string {
	var ID string
	if value, ok := c.Attrs["id"]; ok {
		ID = fmt.Sprint(value)
	} else {
		return "id_" + c.Name
	}

	return ID
}

func (c CharField) LabelTag() string {
	var labelText string
	if c.Label != "" {
		labelText = c.Label
	} else if c.Name != "" {
		labelText = Humanize(c.Name)
	}
	ctx := pongo2.Context{
		"use_tag": true,
		"tag":     "label",
		"label":   labelText + ": ",
		"widget": map[string]any{
			"attrs": map[string]any{
				"for": c.getID(),
			},
		},
	}
	name := "label.html"
	tpl, err := renderTPL(name, ctx)
	if err != nil {
		panic(err)
	}
	return tpl
}

// password

type PasswordField struct {
	BaseField
	Value any `validators:"required"`
}

// email

type EmailField struct {
	BaseField
}

// number
type NumberField struct {
	BaseField
}

// tel
type TelField struct {
	BaseField
}

// url

type URLField struct {
	BaseField
}

// search
type SearchField struct {
	BaseField
}

// date

type DateField struct {
	BaseField
}

// time

type TimeField struct {
	BaseField
}

// datetime-local

type DateTimeLocalField struct {
	BaseField
}

// month

type MonthField struct {
	BaseField
}

// week

type WeekField struct {
	BaseField
}

// color

type ColorField struct {
	BaseField
}

// range

type RangeField struct {
	BaseField
}

// checkbox

type CheckboxField struct {
	BaseField
}

// radio

type RadioField struct {
	BaseField
}

// file

type FileField struct {
	BaseField
}

// hidden

type HiddenField struct {
	BaseField
}

// image
type ImageField struct {
	BaseField
}

// textarea
type TextareaField struct {
	BaseField
}

// select
type SelectField struct {
	BaseField
}

// option

type OptionField struct {
	BaseField
}

// optgroup
type OptionGroupField struct {
	BaseField
}

// datalist

// progress
// legend

func Start(f Field) {

}
