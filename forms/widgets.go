package forms

import (
	"github.com/flosch/pongo2/v6"
)

type Widget interface {
	getContext() map[string]any
	render() string
	getName() string
}

type TextField struct {
	FieldType  string         `json:"type"`
	FieldName  string         `json:"name"`
	FieldValue any            `json:"value"`
	FieldAttrs map[string]any `json:"attr"`
}

func (t TextField) getContext() map[string]any {

	ctx := map[string]any{
		"type":  t.FieldType,
		"name":  t.FieldName,
		"value": t.FieldValue,
		"attrs": t.FieldAttrs,
	}

	return ctx
}

func (t TextField) getName() string {
	return t.FieldName
}

func (t TextField) render() string {
	tpl, err := pongo2.FromFile("forms/templates/widgets/input.html")
	if err != nil {
		panic(err)
	}
	ctx := t.getContext()
	// fmt.Println(ctx)
	uhtml, err := tpl.Execute(pongo2.Context{
		"widget": ctx,
	})

	if err != nil {
		panic(err)
	}
	return uhtml
}

func (t TextField) String() string {
	return t.render()
}
func (t TextField) Css_classes() string {
	return "mx-auto"
}

// type DateField struct {
// 	TextField
// }

func Render(w Widget) string {
	return w.render()
}
