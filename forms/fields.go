package forms

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type Field interface {
	getContext() pongo2.Context
	render() string
	LabelTag() *pongo2.Value
	getID() string
	GetValue() any
	CssClasses() string
}

type BaseField struct {
}

// text

type CharField struct {
	BaseField
	InputType   string
	Name        string
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

func (c CharField) getInputType() string {
	var inpType string = "text"
	if c.InputType != "" {
		inpType = c.InputType
	}
	return inpType
}

func (c CharField) CssClasses() string {
	if value, ok := c.Attrs["class"]; ok {
		return fmt.Sprint(value)
	} else {
		return ""
	}
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
	return c.render()
}

func (c CharField) render() string {
	ctx := c.getContext()
	name := "widgets/input.html"
	tpl, err := renderTPL(name, ctx)
	if err != nil {
		panic(err)
	}
	return tpl
}

func (c CharField) GetValue() any {
	if c.Value != "" {
		value := string(c.Value)
		return value
	}
	return c.EmptyValue
}

func (c CharField) getContext() pongo2.Context {

	ctx := pongo2.Context{
		"widget": map[string]any{
			"type":             c.getInputType(),
			"help_text":        c.HelpText,
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

func (c CharField) GetLabel() string {
	var labelText string
	if c.Label != "" {
		labelText = c.Label
	} else if c.Name != "" {
		labelText = Humanize(c.Name)
	}
	return labelText
}

func (c CharField) LabelTag() *pongo2.Value {
	var labelText string
	if c.Label != "" {
		labelText = c.Label
	} else if c.Name != "" {
		labelText = Humanize(c.Name)
	}
	ctx := pongo2.Context{
		"use_tag": true,
		"tag":     "label",
		"label":   labelText + " ",
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
	return pongo2.AsSafeValue(tpl)
}
