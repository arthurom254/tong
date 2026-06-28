package forms

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type Form struct {
	Fields []Widget
	Errors *map[string]any
	Data   any
}

func (f Form) getContext() pongo2.Context {
	ctx := pongo2.Context{
		"errors":        f.Errors,
		"fields":        f.Fields,
		"hidden_fields": []Widget{},
	}
	return ctx
}

func (f Form) render() string {
	tpl, err := pongo2.FromFile("forms/templates/div.html")
	if err != nil {
		panic(err)
	}
	ctx := f.getContext()
	// fmt.Println(ctx)
	uhtml, err := tpl.Execute(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println(uhtml)
	return uhtml
}
func (f Form) AsP() string {
	return "<p>Paragrap</p>"
}

func (f Form) AsDiv() string {

	return f.render()
}

func (f Form) IsBound() bool {
	return f.Data != nil
}

func (f Form) String() string {
	return f.AsDiv()
}

func (f *Form) Clean() error {
	if !f.IsBound() {
		return nil
	}
	for i := 0; i < len(f.Fields); i++ {
		fmt.Println("Validated Form")
	}
	return nil
}

func (f Form) GetNames() []string {

	fields := []string{}

	for i := 0; i < len(f.Fields); i++ {
		fields = append(fields, f.Fields[i].getName())
	}

	return fields

}
